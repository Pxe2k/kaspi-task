package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Pxe2k/kaspi-task/internal/domain"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBConn interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Repository struct {
	db DBConn
}

func New(db DBConn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get(ctx context.Context, docID string) (domain.Person, error) {
	q := `select * from persons where document_id = $1;`

	rows, err := r.db.Query(ctx, q, docID)
	if err != nil {
		return domain.Person{}, fmt.Errorf("query: %w", err)
	}

	var p domain.Person
	if err = pgxscan.ScanOne(&p, rows); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return domain.Person{}, fmt.Errorf("scan: %w", err)
		}

		return domain.Person{}, fmt.Errorf("person not found with document id: %s", docID)
	}

	return p, nil
}

func (r *Repository) Find(ctx context.Context, name string) ([]domain.Person, error) {
	q := `select * from persons where name ILIKE $1;`

	rows, err := r.db.Query(ctx, q, "%"+name+"%")
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	p := make([]domain.Person, 0)
	if err = pgxscan.ScanAll(&p, rows); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("scan: %w", err)
		}

		return nil, fmt.Errorf("persons not found")
	}

	return p, nil
}

func (r *Repository) Store(ctx context.Context, person domain.Person) error {
	q := `insert into persons (document_id, name, phone) 
			values ($1, $2, $3);`
	_, err := r.db.Exec(ctx, q,
		person.DocumentID,
		person.Name,
		person.Phone,
	)
	if err != nil {
		return fmt.Errorf("insert person: %w", err)
	}

	return nil
}
