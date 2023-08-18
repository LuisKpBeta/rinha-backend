package repository

import (
	"database/sql"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/google/uuid"
)

type PeopleRepository struct {
	DbConn *sql.DB
}

func CreatePeopleRepository(db *sql.DB) people.CreatePeopleRepository {
	return &PeopleRepository{
		DbConn: db,
	}
}
func CreateFindPeopleRepository(db *sql.DB) people.FindPeopleByIdRepository {
	return &PeopleRepository{
		DbConn: db,
	}
}

func (p *PeopleRepository) NickNameExists(nickname string) (bool, error) {
	var id string
	err := p.DbConn.QueryRow("SELECT id FROM people where nickname = $1", nickname).
		Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *PeopleRepository) Create(people *people.People) error {
	people.Id = uuid.NewString()
	stmt, err := p.DbConn.Prepare("INSERT INTO people (id, nickname, name, birthdate, stack)  VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.QueryRow(people.Id, people.Nickname, people.Name, people.Birthday, people.Stacks)
	return nil
}

func (p *PeopleRepository) Find(id uuid.UUID) (*people.People, error) {
	var people people.People
	err := p.DbConn.QueryRow("SELECT id, nickname, name, birthdate, stack FROM people WHERE id = $1", id).
		Scan(&people.Id, &people.Nickname, &people.Name, &people.Birthday, &people.Stacks)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &people, nil
}
