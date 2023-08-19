package test_people

import (
	"net/http"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/goccy/go-json"
)

func (suite *TaskApiTestSuite) Test_SearchPeopleWithEmptyTerm() {
	url := "/pessoas?t="
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusBadRequest, res.StatusCode)
}
func (suite *TaskApiTestSuite) Test_SearchPeopleWithNoResult() {
	url := "/pessoas?t=anyone"
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusOK, res.StatusCode)
	peopleRes := []PeopleResponse{}
	derr := json.NewDecoder(res.Body).Decode(&peopleRes)
	if derr != nil {
		panic(derr)
	}

	suite.Equal(0, len(peopleRes))
}
func (suite *TaskApiTestSuite) Test_SearchPeopleWithMatchName() {
	people1 := people.People{
		Nickname: "already_exists",
		Name:     "dummy",
		Birthday: "1999-01-01",
		Stacks:   "C#, javascript",
	}
	people2 := people.People{
		Nickname: "already_exists",
		Name:     "other name",
		Birthday: "1999-01-01",
		Stacks:   "C#, javascript",
	}

	suite.InsertPeopleDb(&people1)
	suite.InsertPeopleDb(&people2)
	url := "/pessoas?t=umm"
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	testCheck := suite.Equal(http.StatusOK, res.StatusCode)
	if !testCheck {
		return
	}
	peopleRes := []PeopleResponse{}
	derr := json.NewDecoder(res.Body).Decode(&peopleRes)
	if derr != nil {
		panic(derr)
	}

	lenOk := suite.Equal(1, len(peopleRes))
	if !lenOk {
		return
	}
	suite.Equal(people1.Id, peopleRes[0].Id)
	suite.Equal(people1.Name, peopleRes[0].Nome)
}
func (suite *TaskApiTestSuite) Test_SearchPeopleWithMatchNickname() {
	people1 := people.People{
		Nickname: "jhondoe",
		Name:     "first",
		Birthday: "1999-01-01",
	}
	people2 := people.People{
		Nickname: "nickwithoutname",
		Name:     "name",
		Birthday: "1999-01-01",
	}

	suite.InsertPeopleDb(&people1)
	suite.InsertPeopleDb(&people2)
	url := "/pessoas?t=ond"
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	testCheck := suite.Equal(http.StatusOK, res.StatusCode)
	if !testCheck {
		return
	}
	peopleRes := []PeopleResponse{}
	derr := json.NewDecoder(res.Body).Decode(&peopleRes)
	if derr != nil {
		panic(derr)
	}

	lenOk := suite.Equal(1, len(peopleRes))
	if !lenOk {
		return
	}
	suite.Equal(people1.Id, peopleRes[0].Id)
	suite.Equal(people1.Name, peopleRes[0].Nome)
}
func (suite *TaskApiTestSuite) Test_SearchPeopleWithMatchNameAndNickName() {
	people1 := people.People{
		Nickname: "jhondoe",
		Name:     "first",
		Birthday: "1999-01-01",
	}
	people2 := people.People{
		Nickname: "nickwithoutname",
		Name:     "second",
		Birthday: "1999-01-01",
	}

	suite.InsertPeopleDb(&people1)
	suite.InsertPeopleDb(&people2)
	url := "/pessoas?t=ond"
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	testCheck := suite.Equal(http.StatusOK, res.StatusCode)
	if !testCheck {
		return
	}
	peopleRes := []PeopleResponse{}
	derr := json.NewDecoder(res.Body).Decode(&peopleRes)
	if derr != nil {
		panic(derr)
	}

	lenOk := suite.Equal(2, len(peopleRes))
	if !lenOk {
		return
	}
	suite.Equal(people1.Id, peopleRes[0].Id)
	suite.Equal(people1.Name, peopleRes[0].Nome)
	suite.Equal(people2.Id, peopleRes[1].Id)
	suite.Equal(people2.Name, peopleRes[1].Nome)
}

func (suite *TaskApiTestSuite) Test_SearchPeopleWithMatchStackAndName() {
	people1 := people.People{
		Nickname: "jhondeno",
		Name:     "first",
		Birthday: "1999-01-01",
	}
	people2 := people.People{
		Nickname: "nickwithoutname",
		Name:     "second",
		Birthday: "1999-01-01",
		Stacks:   "Deno, Go",
	}

	suite.InsertPeopleDb(&people1)
	suite.InsertPeopleDb(&people2)
	url := "/pessoas?t=ond"
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	testCheck := suite.Equal(http.StatusOK, res.StatusCode)
	if !testCheck {
		return
	}
	peopleRes := []PeopleResponse{}
	derr := json.NewDecoder(res.Body).Decode(&peopleRes)
	if derr != nil {
		panic(derr)
	}

	lenOk := suite.Equal(2, len(peopleRes))
	if !lenOk {
		return
	}
	suite.Equal(people1.Id, peopleRes[0].Id)
	suite.Equal(people1.Name, peopleRes[0].Nome)
	suite.Equal(people2.Id, peopleRes[1].Id)
	suite.Equal(people2.Name, peopleRes[1].Nome)
}
