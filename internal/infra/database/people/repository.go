package repository

import (
	"context"
	"time"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/rueidis"
)

const (
	MAX_SEARCH_RESULT = 50
)

type PeopleRepository struct {
	DbConn   *pgxpool.Pool
	CacheCon rueidis.Client
}

func CreatePeopleRepository(db *pgxpool.Pool, cache rueidis.Client) people.CreatePeopleRepository {
	return &PeopleRepository{
		DbConn:   db,
		CacheCon: cache,
	}
}
func CreateFindPeopleRepository(db *pgxpool.Pool, cache rueidis.Client) people.FindPeopleByIdRepository {
	return &PeopleRepository{
		DbConn:   db,
		CacheCon: cache,
	}
}
func CreateSearchPeopleRepository(db *pgxpool.Pool, cache rueidis.Client) people.SearchPeopleRepository {
	return &PeopleRepository{
		DbConn:   db,
		CacheCon: cache,
	}
}
func CreateCountPeopleRepository(db *pgxpool.Pool, cache rueidis.Client) people.CountPeopleRepository {
	return &PeopleRepository{
		DbConn:   db,
		CacheCon: cache,
	}
}

func (p *PeopleRepository) NickNameExists(nickname string) (bool, error) {
	exists, _ := p.CacheCheckNicknameExists(nickname)
	if exists {
		return true, nil
	}

	var id string
	err := p.DbConn.QueryRow(context.Background(), "SELECT id FROM people where nickname = $1", nickname).
		Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *PeopleRepository) Create(people *people.People) error {
	people.Id = uuid.NewString()
	_, err := p.DbConn.Exec(
		context.Background(),
		"INSERT INTO people (id, nickname, name, birthday, stacks)  VALUES ($1, $2, $3, $4, $5)",
		people.Id, people.Nickname, people.Name, people.Birthday, people.Stacks)
	if err != nil {
		return err
	}
	go p.CacheSavePeople(people)

	return nil
}

func (p *PeopleRepository) Find(id uuid.UUID) (*people.People, error) {
	peoplecache, _ := p.CacheCheckIdExists(id)
	if peoplecache != nil {
		return peoplecache, nil
	}
	var people people.People

	var birthday time.Time
	err := p.DbConn.QueryRow(context.Background(), "SELECT id, nickname, name, birthday, stacks FROM people WHERE id = $1", id).
		Scan(&people.Id, &people.Nickname, &people.Name, &birthday, &people.Stacks)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	people.Birthday = birthday.Format("2006-01-02")
	return &people, nil
}
func (p *PeopleRepository) SearchPeople(term string) ([]people.People, error) {
	queryTerm := "%" + term + "%"
	rows, err := p.DbConn.Query(context.Background(), "SELECT id, name, nickname, birthday, stacks FROM people WHERE nickname ILIKE $1 or name ILIKE $1 or stacks ILIKE $1 LIMIT $2", queryTerm, MAX_SEARCH_RESULT)
	if err != nil {
		return nil, err
	}

	var peopleList []people.People

	for rows.Next() {
		p := people.People{}
		var birthday time.Time
		err := rows.Scan(&p.Id, &p.Name, &p.Nickname, &birthday, &p.Stacks)
		if err != nil {
			return nil, err
		}
		p.Birthday = birthday.Format("2006-01-02")
		peopleList = append(peopleList, p)
	}
	return peopleList, nil

}

func (p *PeopleRepository) Count() (int, error) {
	var total int
	err := p.DbConn.QueryRow(context.Background(), "SELECT Count(id) FROM people").Scan(&total)

	if err != nil {
		return 0, err
	}
	return total, nil
}
