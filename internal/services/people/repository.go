package people

type CreatePeopleRepository interface {
	Create(people *People) error
	NickNameExists(name string) (bool, error)
}
