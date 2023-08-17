package test_create_people

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	pc "github.com/LuisKpBeta/rinha-backend/internal/infra/api/controllers"
	"github.com/LuisKpBeta/rinha-backend/internal/infra/database"
	repo "github.com/LuisKpBeta/rinha-backend/internal/infra/database/people"
	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type TaskApiTestSuite struct {
	suite.Suite
	Db *sql.DB
	r  *gin.Engine
	ts *httptest.Server
}

func (suite *TaskApiTestSuite) SetupSuite() {
	suite.r = gin.Default()
	suite.Db = database.ConnectToDatabase()
	suite.SetupHttpServer()
	suite.RunHttpServer()
}
func (suite *TaskApiTestSuite) TearDownSuite() {
	suite.Db.Close()
}
func (suite *TaskApiTestSuite) SetupTest() {
	stmt, _ := suite.Db.Prepare("DELETE FROM people")
	stmt.Exec()
}
func (suite *TaskApiTestSuite) SetupHttpServer() {
	peopleRepo := repo.CreatePeopleRepository(suite.Db)
	create := &people.CreatePeople{
		Repository: peopleRepo,
	}
	controller := pc.CreatePeopleController(create)

	suite.r.POST("/people", controller.Create)
}
func (suite *TaskApiTestSuite) RunHttpServer() {
	suite.ts = httptest.NewServer(suite.r)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(TaskApiTestSuite))
}

type CreatePeoplePayload struct {
	Apelido    string   `json:"apelido,omitempty"`
	Nome       string   `json:"nome,omitempty"`
	Nascimento string   `json:"nascimento,omitempty"`
	Stacks     []string `json:"stacks"`
}
type CreatePeoplePayloadInvalid struct {
	Apelido    interface{}   `json:"apelido,omitempty"`
	Nome       string        `json:"nome,omitempty"`
	Nascimento string        `json:"nascimento,omitempty"`
	Stacks     []interface{} `json:"stacks"`
}
type ErrorResponse struct {
	Error string
}

func (suite *TaskApiTestSuite) InsertPeopleDb(p *people.People) {
	p.Id = uuid.NewString()
	stmt, err := suite.Db.Prepare("INSERT INTO people (id, nickname, name, birthdate, stack)  VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	stmt.QueryRow(p.Id, p.Nickname, p.Name, p.Birthday, p.Stacks)
}

func (suite *TaskApiTestSuite) makeHttpPost(payload interface{}, url string) (*http.Response, error) {
	jsonData, _ := json.Marshal(payload)
	client := &http.Client{}
	request, _ := http.NewRequest("POST", suite.ts.URL+url, bytes.NewBuffer(jsonData))
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
	res, err := suite.makeHttpPost(payload, "/people")
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
	res, err := suite.makeHttpPost(payload, "/people")
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
	res, err := suite.makeHttpPost(payload, "/people")
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
	res, err := suite.makeHttpPost(payload, "/people")
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
	res, err := suite.makeHttpPost(payload, "/people")
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
	res, err := suite.makeHttpPost(payload, "/people")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusBadRequest, res.StatusCode)
}
func (suite *TaskApiTestSuite) Test_CreatePeopleWithError_StackHasNumber() {
	payload := CreatePeoplePayloadInvalid{
		Apelido:    "dummy",
		Nascimento: "1999-01-01",
		Stacks:     []interface{}{1, "stack"},
	}
	res, err := suite.makeHttpPost(payload, "/people")
	suite.NoError(err)
	defer res.Body.Close()
	suite.Equal(http.StatusBadRequest, res.StatusCode)
}
