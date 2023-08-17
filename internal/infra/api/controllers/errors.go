package people_controller

import "errors"

var (
	ErrTooLongNickName  = errors.New("apelido deve ter no maximo 32 characteres")
	ErrEmptyNickName    = errors.New("apelido não deve estar vazio")
	ErrTooLongName      = errors.New("nome deve ter no maximo 100 characteres")
	ErrEmptyName        = errors.New("nome não deve estar vazio")
	ErrTooLongStackName = errors.New("stack trabalhada deve ter no maximo 32 characteres")
	ErrInvalidBirthday  = errors.New("nascimento deve seguir o formato AAAA-MM-DD")
)
