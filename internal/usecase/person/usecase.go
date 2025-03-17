package person

type UseCase struct {
	repository repository
}

func New(r repository) *UseCase {
	return &UseCase{
		repository: r,
	}
}
