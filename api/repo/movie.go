package repo

import (
	"movie_planet/infra"
	"movie_planet/models"
)

type MovieRepo struct {
	db infra.Database
}

func NewMovieRepo(db infra.Database) MovieRepo {
	return MovieRepo{
		db: db,
	}
}

func (m MovieRepo) FindAll(movie models.Movie, keyword string) (*[]models.Movie, int64, error) {
	var movies []models.Movie
	var totalRows int64 = 0

	queryBuilder := m.db.DB.Model(&models.Movie{})

	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(m.db.DB.Where("movie.name LIKE ? ", queryKeyword))
	}

	err := queryBuilder.Where(movie).Find(&movies).Count(&totalRows).Error

	return &movies, totalRows, err
}

func (m MovieRepo) Find(movie models.Movie) (models.Movie, error) {
	var movies models.Movie
	err := m.db.DB.Model(&models.Movie{}).Where(&movie).Take(&movies).Error
	return movies, err
}
func (m MovieRepo) Create(movie models.Movie) error {
	return m.db.DB.Create(&movie).Error
}

func (m MovieRepo) Update(movie models.Movie) error {
	return m.db.DB.Save(&movie).Error
}

func (m MovieRepo) Delete(movie models.Movie) error {
	return m.db.DB.Delete(movie).Error
}
