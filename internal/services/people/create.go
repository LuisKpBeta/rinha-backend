package people

type CreatePeople struct {
	Repository CreatePeopleRepository
}

func (c *CreatePeople) Create(nickname string, name string, birthday string, stack []string) (*People, error) {
	exists, err := c.Repository.NickNameExists(nickname)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrNickNameAlreadyExists
	}

	newPeople := People{
		Nickname: nickname,
		Name:     name,
		Birthday: birthday,
	}
	newPeople.SetStacksFromArray(stack)
	err = c.Repository.Create(&newPeople)
	if err != nil {
		return nil, nil
	}
	return &newPeople, nil
}
