package test_people

import (
	"net/http"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

func (suite *TaskApiTestSuite) Test_FindPeopleByIdWhenExists() {
	people := people.People{
		Nickname: "already_exists",
		Name:     "dummy",
		Birthday: "1999-01-01",
		Stacks:   "C#, javascript",
	}
	suite.InsertPeopleDb(&people)
	url := "/people/" + people.Id
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	testCheck := suite.Equal(http.StatusOK, res.StatusCode)
	if !testCheck {
		return
	}
	peopleRes := &PeopleResponse{}
	derr := json.NewDecoder(res.Body).Decode(peopleRes)
	if derr != nil {
		panic(derr)
	}
	suite.Equal(people.Id, peopleRes.Id)
	suite.Equal(people.Name, peopleRes.Nome)
	suite.Equal(people.Nickname, peopleRes.Apelido)
	suite.Equal(people.Birthday, peopleRes.Nascimento)
	suite.Equal(len(people.GetArrayFromStringStack()), len(peopleRes.Stacks))
}
func (suite *TaskApiTestSuite) Test_FindPeopleByIdWhenNotExists() {
	url := "/people/" + uuid.NewString()
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusNotFound, res.StatusCode)

}
