package test_people

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
)

func (suite *TaskApiTestSuite) Test_CountTotalOfPeople() {
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
	url := "/contagem-pessoas"
	res, err := suite.makeHttpGet(url)
	suite.NoError(err)
	defer res.Body.Close()
	testCheck := suite.Equal(http.StatusOK, res.StatusCode)
	if !testCheck {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	i, err := strconv.Atoi(string(body))
	if err != nil {
		panic(err)
	}

	suite.Equal(i, 2)

}
