package repository

import (
	"database/sql"

	"github.com/ReyLegar/vkTestProject/internal/models"
)

type MoviePostgres struct {
	db *sql.DB
}

func NewMoviePostgres(db *sql.DB) *MoviePostgres {
	return &MoviePostgres{db: db}

}

func (m *MoviePostgres) Create(movie models.Movie) (int, error) {
	return 5, nil
}

func (m *MoviePostgres) Update(movieID int, movie models.Movie) error {
	return nil
}

func (m *MoviePostgres) Delete(movieID int) error {
	return nil
}

func (m *MoviePostgres) GetAll() ([]models.Movie, error) {
	return nil, nil
}

func (m *MoviePostgres) SearchByTitle(title string) ([]models.Movie, error) {
	return nil, nil
}

func (m *MoviePostgres) SearchByActorName(actorName string) ([]models.Movie, error) {
	return nil, nil
}
