package people

type SearchPeople struct {
	Repository SearchPeopleRepository
}

func (s *SearchPeople) SearchByTerm(term string) ([]People, error) {
	people, err := s.Repository.SearchPeople(term)
	return people, err
}
