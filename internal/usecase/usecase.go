package usecase

type UseCase struct {
	repository repository
}

func New(r repository) *UseCase {
	return &UseCase{
		repository: r,
	}
}
