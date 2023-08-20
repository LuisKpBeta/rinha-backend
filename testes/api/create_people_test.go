package test_people

import (
	"bytes"
	"context"
	"net/http"

	pc "github.com/LuisKpBeta/rinha-backend/internal/infra/api/controllers"
	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

func (suite *TaskApiTestSuite) InsertPeopleDb(p *people.People) {
	p.Id = uuid.NewString()
	suite.Db.Exec(context.Background(), "INSERT INTO people (id, nickname, name, birthday, stacks)  VALUES ($1, $2, $3, $4, $5)", p.Id, p.Nickname, p.Name, p.Birthday, p.Stacks)

}

func (suite *TaskApiTestSuite) makeHttpPost(payload interface{}, url string) (*http.Response, error) {
	jsonData, _ := json.Marshal(payload)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", suite.ts.URL+url, bytes.NewBuffer(jsonData))
	request.Header.Add("Content-Type", "application/json")
	res, err := client.Do(request)

	return res, err
}
func (suite *TaskApiTestSuite) makeHttpGet(url string) (*http.Response, error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", suite.ts.URL+url, nil)
	request.Header.Add("Content-Type", "application/json")
	res, err := client.Do(request)

	return res, err
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithSuccess() {
	payload := CreatePeoplePayload{
		Apelido:    "apelido_example",
		Nome:       "nome_example",
		Nascimento: "1999-01-01",
		Stacks:     []string{"stack1", "stack2", "stack3"},
	}
	res, err := suite.makeHttpPost(payload, "/pessoas")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusCreated, res.StatusCode)
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithSuccess_EmptyStacks() {
	payload := CreatePeoplePayload{
		Apelido:    "apelido_example",
		Nome:       "nome_example",
		Nascimento: "1999-01-01",
		Stacks:     nil,
	}
	res, err := suite.makeHttpPost(payload, "/pessoas")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusCreated, res.StatusCode)
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithError_NameAlreadyExists() {
	people := people.People{
		Nickname: "already_exists",
		Name:     "dummy",
		Birthday: "1999-01-01",
	}
	suite.InsertPeopleDb(&people)
	payload := CreatePeoplePayload{
		Apelido:    "already_exists",
		Nome:       "nome_example",
		Nascimento: "1999-01-01",
		Stacks:     nil,
	}
	res, err := suite.makeHttpPost(payload, "/pessoas")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusUnprocessableEntity, res.StatusCode)
	errResponse := &ErrorResponse{}
	derr := json.NewDecoder(res.Body).Decode(errResponse)
	if derr != nil {
		panic(derr)
	}
	suite.Equal(errResponse.Error, "apelido já existe")
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithError_NameIsNull() {
	payload := CreatePeoplePayload{
		Apelido:    "dummy",
		Nascimento: "1999-01-01",
	}
	res, err := suite.makeHttpPost(payload, "/pessoas")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusUnprocessableEntity, res.StatusCode)
	errResponse := &ErrorResponse{}
	derr := json.NewDecoder(res.Body).Decode(errResponse)
	if derr != nil {
		panic(derr)
	}
	suite.Equal(errResponse.Error, pc.ErrEmptyName.Error())
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithError_NickNameIsNull() {
	payload := CreatePeoplePayload{
		Nome:       "dummy",
		Nascimento: "1999-01-01",
	}
	res, err := suite.makeHttpPost(payload, "/pessoas")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusUnprocessableEntity, res.StatusCode)
	errResponse := &ErrorResponse{}
	derr := json.NewDecoder(res.Body).Decode(errResponse)
	if derr != nil {
		panic(derr)
	}
	suite.Equal(errResponse.Error, pc.ErrEmptyNickName.Error())
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithError_NameIsOnlyNumber() {
	payload := CreatePeoplePayloadInvalid{
		Apelido:    1,
		Nascimento: "1999-01-01",
	}
	res, err := suite.makeHttpPost(payload, "/pessoas")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusBadRequest, res.StatusCode)
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithError_StackHasNumber() {
	payload := CreatePeoplePayloadInvalid{
		Apelido:    "dummy",
		Nascimento: "1999-01-01",
		Stacks:     []interface{}{1, "stacks"},
	}
	res, err := suite.makeHttpPost(payload, "/pessoas")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusBadRequest, res.StatusCode)
}
