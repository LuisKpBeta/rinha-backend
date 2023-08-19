package repository

import (
	"database/sql"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/google/uuid"
)

const (
	MAX_SEARCH_RESULT = 50
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
func CreateSearchPeopleRepository(db *sql.DB) people.SearchPeopleRepository {
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
	stmt, err := p.DbConn.Prepare("INSERT INTO people (id, nickname, name, birthday, stacks)  VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.QueryRow(people.Id, people.Nickname, people.Name, people.Birthday, people.Stacks)
	return nil
}

func (p *PeopleRepository) Find(id uuid.UUID) (*people.People, error) {
	var people people.People
	err := p.DbConn.QueryRow("SELECT id, nickname, name, birthday, stacks FROM people WHERE id = $1", id).
		Scan(&people.Id, &people.Nickname, &people.Name, &people.Birthday, &people.Stacks)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &people, nil
}
func (p *PeopleRepository) SearchPeople(term string) ([]people.People, error) {

	stmt, err := p.DbConn.Prepare("SELECT id, name, nickname, birthday, stacks FROM people WHERE nickname ILIKE $1 or name ILIKE $1 or stacks ILIKE $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	queryTerm := "%" + term + "%"
	rows, err := stmt.Query(queryTerm)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	i := 0
	var peopleList []people.People

	for rows.Next() && i < MAX_SEARCH_RESULT {
		p := people.People{}
		err := rows.Scan(&p.Id, &p.Name, &p.Nickname, &p.Birthday, &p.Stacks)
		if err != nil {
			return nil, err
		}
		peopleList = append(peopleList, p)
		i++
	}
	return peopleList, nil

}
