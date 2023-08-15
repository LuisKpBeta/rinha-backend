package people

type CreatePeopleRepository interface {
	Create(people *People) (*People, error)
}
