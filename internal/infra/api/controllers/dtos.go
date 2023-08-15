package people_controller

type CreatePopleDto struct {
	Nickname string   `json:"apelido" validate:"required,max=32"`
	Name     string   `json:"nome" validate:"required,max=100"`
	Birthday string   `json:"nascimento" validate:"required,datetime=2006-01-02"`
	Stacks   []string `json:"stacks" validate:"dive, max=32"`
}
