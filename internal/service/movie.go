package service

import (
	"github.com/ReyLegar/vkTestProject/internal/models"
	"github.com/ReyLegar/vkTestProject/internal/repository"
)

type MovieService struct {
	movieRepo repository.MovieRepository
	actorRepo repository.ActorRepository
}

func NewMovieService(movieRepo repository.MovieRepository, actorRepo repository.ActorRepository) *MovieService {
	return &MovieService{movieRepo: movieRepo, actorRepo: actorRepo}
}

func (s *MovieService) AddMovie(movie models.Movie) (int, error) {

	for _, actor := range movie.Actors {
		_, err := s.actorRepo.GetByName(actor.Name)
		if err != nil {
			_, err := s.actorRepo.Create(actor)
			if err != nil {
				return 0, err
			}
		}
	}

	movieID, err := s.movieRepo.Create(movie)
	if err != nil {
		return 0, err
	}

	return movieID, nil
}

func (s *MovieService) UpdateMovie(movieID int, movie models.Movie) error {

	for _, actor := range movie.Actors {
		_, err := s.actorRepo.GetByName(actor.Name)
		if err != nil {
			_, err := s.actorRepo.Create(actor)
			if err != nil {
				return err
			}
		}
	}

	err := s.movieRepo.Update(movieID, movie)
	if err != nil {
		return err
	}

	return nil
}

func (s *MovieService) DeleteMovie(movieID int) error {
	err := s.movieRepo.Delete(movieID)
	if err != nil {
		return err
	}

	return nil
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	return s.movieRepo.GetAll()
}

func (s *MovieService) SearchMoviesByTitle(title string) ([]models.Movie, error) {
	return s.movieRepo.SearchByTitle(title)
}

func (s *MovieService) SearchMoviesByActorName(actorName string) ([]models.Movie, error) {
	return s.movieRepo.SearchByActorName(actorName)
}
