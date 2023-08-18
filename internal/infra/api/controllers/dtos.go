package people_controller

import (
	"time"
)

var (
	DATE_LAYOUT = "2006-01-02"
)

type CreatePopleDto struct {
	Nickname string   `json:"apelido" validate:"required,max=32"`
	Name     string   `json:"nome" validate:"required,max=100"`
	Birthday string   `json:"nascimento" validate:"required,datetime=2006-01-02"`
	Stacks   []string `json:"stacks" validate:"dive, max=32"`
}
type ReadPopleDto struct {
	Id       string   `json:"id"`
	Nickname string   `json:"apelido" `
	Name     string   `json:"nome" `
	Birthday string   `json:"nascimento" `
	Stacks   []string `json:"stacks" `
}

func (c *CreatePopleDto) IsValid() error {
	if len(c.Nickname) == 0 {
		return ErrEmptyNickName
	}
	if len(c.Nickname) > 32 {
		return ErrTooLongNickName
	}
	if len(c.Name) == 0 {
		return ErrEmptyName
	}
	if len(c.Name) > 100 {
		return ErrTooLongName
	}
	if _, err := time.Parse(DATE_LAYOUT, c.Birthday); err != nil {
		return ErrInvalidBirthday
	}
	for _, st := range c.Stacks {
		if len(st) > 32 {
			return ErrTooLongStackName
		}
	}
	return nil
}
