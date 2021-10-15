package repo

import (
	"movie_planet/infra"
	"movie_planet/models"
)

type RateRepo struct {
	db infra.Database
}

func NewRateRepo(db infra.Database) RateRepo {
	return RateRepo{
		db: db,
	}
}

func (r RateRepo) FindAll(rate models.Rate, movieID int64) (*[]models.Rate, int64, error) {
	var rates []models.Rate
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Model(&models.Rate{})
	err := queryBuilder.Where("movie_id = ?", movieID).Find(&rates).Count(&totalRows).Error
	return &rates, totalRows, err
}

func (r RateRepo) Create(rate models.Rate) error {
	var newRate models.Rate
	newRate.Comment = rate.Comment
	newRate.MovieID = rate.MovieID
	newRate.UserID = rate.UserID
	newRate.Score = rate.Score
	newRate.Status = rate.Status
	return r.db.DB.Create(&newRate).Error
}
