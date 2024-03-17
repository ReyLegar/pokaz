package repository

import (
	"database/sql"
	"fmt"

	"time"

	"github.com/ReyLegar/vkTestProject/internal/models"
	_ "github.com/lib/pq"
)

type ActorPostgres struct {
	db *sql.DB
}

func NewActorPostgres(db *sql.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (a *ActorPostgres) Create(actor models.Actor) (int, error) {
	query := `INSERT INTO Actors (Name, Gender, BirthDate) VALUES ($1, $2, $3) RETURNING ActorID`
	var actorID int
	date, err := time.Parse("2006-01-02", actor.BirthDate)

	if err != nil {
		fmt.Println(err)
	}

	err = a.db.QueryRow(query, actor.Name, actor.Gender, date).Scan(&actorID)
	if err != nil {
		return 0, err
	}
	return actorID, nil
}

func (a *ActorPostgres) Update(actorID int, actor models.Actor) error {
	query := `
		UPDATE Actors
		SET Name = $2, Gender = $3, BirthDate = $4
		WHERE ActorID = $1
	`

	date, err := time.Parse("2006-01-02", actor.BirthDate)
	if err != nil {
		return err
	}

	_, err = a.db.Exec(query, actorID, actor.Name, actor.Gender, date)
	if err != nil {
		return err
	}

	return nil
}

func (a *ActorPostgres) Delete(actorID int) error {
	query := `
		DELETE FROM Actors
		WHERE ActorID = $1
	`

	_, err := a.db.Exec(query, actorID)
	if err != nil {
		return err
	}

	return nil
}

func (a *ActorPostgres) GetByName(name string) (models.Actor, error) {
	var actor models.Actor
	query := "SELECT * FROM Actors WHERE Name = $1"
	err := a.db.QueryRow(query, name).Scan(&actor.ActorID, &actor.Name, &actor.Gender, &actor.BirthDate)
	if err != nil {
		return models.Actor{}, err
	}
	return actor, nil
}
