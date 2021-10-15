package service

import (
	"movie_planet/api/repo"
	"movie_planet/models"
)

type RateService struct {
	repo repo.RateRepo
}

func NewRateService(r repo.RateRepo) RateService {
	return RateService{
		repo: r,
	}
}

func (r RateService) Create(rate models.Rate) error {
	return r.repo.Create(rate)
}

func (r RateService) FindAll(rate models.Rate, movieID int64) (*[]models.Rate, int64, error) {
	return r.repo.FindAll(rate, movieID)
}
