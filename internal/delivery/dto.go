package delivery

import (
	"fmt"
	"strings"
)

type StoreRequest struct {
	DocumentID string `json:"document_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}

type StoreResponse struct {
	Status string `json:"status"`
}

func (r *StoreRequest) Validate() error {
	if r.DocumentID == "" || len(r.DocumentID) != 12 {
		return fmt.Errorf("document_id is required and must be 12 characters long")
	}
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.Phone == "" || !strings.HasPrefix(r.Phone, "+") || len(r.Phone) != 11 {
		return fmt.Errorf("phone is required, must start with '+' and be 11 characters long")
	}
	return nil
}

type GetRequest struct {
	DocumentID string `json:"document_id"`
}

type GetResponse struct {
	DocumentID string `json:"document_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}

func (r *GetRequest) validate() error {
	if r.DocumentID == "" && len(r.DocumentID) != 12 {
		return fmt.Errorf("document_id is required and must be 12 characters long")
	}

	return nil
}

type CheckRequest struct {
	DocumentID string `json:"document_id"`
}

type CheckResponse struct {
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

func (r *CheckRequest) validate() error {
	if r.DocumentID == "" && len(r.DocumentID) != 12 {
		return fmt.Errorf("document_id is required and must be 12 characters long")
	}

	return nil
}

type FindRequest struct {
	Name string `json:"name"`
}

type FindResponse struct {
	Persons []personsResponse `json:"persons"`
}

type personsResponse struct {
	DocumentID string `json:"document_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}

func (r *FindRequest) validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}

	return nil
}
