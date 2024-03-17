package service

import (
	"github.com/ReyLegar/vkTestProject/internal/models"
	"github.com/ReyLegar/vkTestProject/internal/repository"
)

type ActorRepository interface {
	AddActor(actor models.Actor) (int, error)
	UpdateActor(actorID int, actor models.Actor) error
	DeleteActor(actorID int) error
}

type MovieRepository interface {
	AddMovie(movie models.Movie) (int, error)
	UpdateMovie(movieID int, movie models.Movie) error
	DeleteMovie(movieID int) error
	GetAllMovies() ([]models.Movie, error)
	SearchMoviesByTitle(title string) ([]models.Movie, error)
	SearchMoviesByActorName(actorName string) ([]models.Movie, error)
}

type Service struct {
	ActorRepository
	MovieRepository
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		ActorRepository: NewActorService(repos.ActorRepository),
		MovieRepository: NewMovieService(repos.MovieRepository, repos.ActorRepository),
	}
}
