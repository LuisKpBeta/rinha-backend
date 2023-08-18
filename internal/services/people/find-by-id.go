package people

import (
	"github.com/google/uuid"
)

type FindPeopleById struct {
	Repository FindPeopleByIdRepository
}

func (f *FindPeopleById) Find(id uuid.UUID) (*People, error) {
	people, err := f.Repository.Find(id)
	return people, err
}
