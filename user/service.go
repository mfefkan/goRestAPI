package user

import (
	"math/rand"
)

type Service interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	UpdateBalance(userID uint, amount int) (*Model, error)
	GuessAndUpdateBalance(userID uint, amount int, targets []int) (map[string]interface{}, error)
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

func (s service) UpdateBalance(userID uint, amount int) (*Model, error) {
	user, err := s.repo.findUserByID(userID)
	if err != nil {
		return nil, err
	}

	user.Balance += amount

	err = s.repo.updateUserBalance(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) GuessAndUpdateBalance(userID uint, amount int, targets []int) (map[string]interface{}, error) {
	user, err := s.repo.findUserByID(userID)
	if err != nil {
		return nil, err
	}

	results := make(map[int]bool)
	randomNumbers := make(map[int]int) // Her hedef için üretilen rastgele sayıları saklamak için
	successful := true

	// Kazanç faktörünü hesapla
	totalGainFactor := s.calculateTotalGainFactor(targets)

	// Her bir hedef için işlem yap
	for _, target := range targets {
		random := rand.Intn(99) + 1
		successfulThisRound := random >= target
		results[target] = successfulThisRound
		randomNumbers[target] = random // Rastgele sayıyı kaydet

		// Başarısız tahmin olduğunda başarı durumunu güncelle
		if !successfulThisRound {
			successful = false
		}
	}

	// Başarılı ise user'ın balance'ına amount*toplam_kazanç_faktörü eklenecek, başarısız ise -amount olacak.
	if successful {
		user.Balance += int(float64(amount) * totalGainFactor)
	} else {
		user.Balance -= amount
	}

	// Veritabanında güncelle
	err = s.repo.updateUserBalance(user)
	if err != nil {
		return nil, err
	}

	// Sonuçları döndür
	result := map[string]interface{}{
		"user":          user,
		"results":       results,
		"randomNumbers": randomNumbers, // Rastgele sayıları da döndür
		"successful":    successful,
	}
	return result, nil
}

func (s service) calculateTotalGainFactor(targets []int) float64 {
	totalGainFactor := 0.0
	for _, target := range targets {
		totalGainFactor += s.calculateGainFactor(target)
	}
	return totalGainFactor
}

func (s service) calculateGainFactor(target int) float64 {
	return 98 / (100 - float64(target))
}
