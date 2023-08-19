package people

type CountPeople struct {
	Repository CountPeopleRepository
}

func (f *CountPeople) Count() (int, error) {
	total, err := f.Repository.Count()
	return total, err
}
