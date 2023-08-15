package people

import "fmt"

type CreatePeople struct {
	repository CreatePeopleRepository
}

func (c *CreatePeople) CreatePeople(nickname string, name string, birthday string, stack []string) error {
	fmt.Println(name, nickname, birthday, stack)
	return nil
}
