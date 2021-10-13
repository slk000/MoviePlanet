package service

import (
	"movie_planet/api/repo"
	"movie_planet/models"
)

type MovieService struct {
	repo repo.MovieRepo
}

func NewMovieService(r repo.MovieRepo) MovieService {
	return MovieService{
		repo: r,
	}
}

func (m MovieService) FindAll(movie models.Movie, keyword string) (*[]models.Movie, int64, error) {
	return m.repo.FindAll(movie, keyword)
}

func (m MovieService) Find(movie models.Movie) (models.Movie, error) {
	return m.repo.Find(movie)
}

func (m MovieService) Create(movie models.Movie) error {
	return m.repo.Create(movie)
}

func (m MovieService) Update(movie models.Movie) error {
	return m.repo.Update(movie)
}

func (m MovieService) Delete(id int64) error {
	var movie models.Movie
	movie.ID = id
	return m.repo.Delete(movie)
}
