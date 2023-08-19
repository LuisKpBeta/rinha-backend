package people

import "github.com/google/uuid"

type CreatePeopleRepository interface {
	Create(people *People) error
	NickNameExists(name string) (bool, error)
}

type FindPeopleByIdRepository interface {
	Find(id uuid.UUID) (*People, error)
}

type SearchPeopleRepository interface {
	SearchPeople(term string) ([]People, error)
}
