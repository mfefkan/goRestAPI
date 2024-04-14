package user

import "gorm.io/gorm"

type Repository interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	findUserByID(userID uint) (*Model, error)
	updateUserBalance(user *Model) error

	Migration() error
}

type repository struct {
	db *gorm.DB
}

var _ Repository = repository{}

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}
}

func (repo repository) Get(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.First(model).Error
	// SELECT * FROM users WHERE id=1
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo repository) Create(model Model) (uint, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (repo repository) findUserByID(userID uint) (*Model, error) {
	var user Model
	result := repo.db.First(&user, userID) // userID'ye göre kullanıcıyı ara
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo repository) updateUserBalance(user *Model) error {
	// user nesnesinde güncellenmiş balance ile
	result := repo.db.Model(user).Update("balance", user.Balance)
	return result.Error
}

func (repo repository) Migration() error {
	return repo.db.AutoMigrate(&Model{})
}
