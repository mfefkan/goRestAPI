package user

type Service interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
}

type service struct {
	repo Repository
}

var _ Service = service{}

func NewService(repo Repository) Service {
	return service{repo: repo}
}

func (s service) Get(id uint) (*Model, error) {
	return s.repo.Get(id)
}

func (s service) Create(model Model) (uint, error) {
	return s.repo.Create(model)
}
