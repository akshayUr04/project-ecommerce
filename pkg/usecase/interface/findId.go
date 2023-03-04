package interfaces

type FindIdUseCase interface {
	FindId(string) (int, error)
}
