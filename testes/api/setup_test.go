package test_people

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	pc "github.com/LuisKpBeta/rinha-backend/internal/infra/api/controllers"
	"github.com/LuisKpBeta/rinha-backend/internal/infra/database"
	repo "github.com/LuisKpBeta/rinha-backend/internal/infra/database/people"
	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/gin-gonic/gin"
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

	findPeopleRepo := repo.CreateFindPeopleRepository(suite.Db)
	findPeople := &people.FindPeopleById{
		Repository: findPeopleRepo,
	}
	searchPeopleRepo := repo.CreateSearchPeopleRepository(suite.Db)
	searchPeople := &people.SearchPeople{
		Repository: searchPeopleRepo,
	}

	controller := pc.CreatePeopleController(create, findPeople, searchPeople)

	suite.r.POST("/people", controller.Create)
	suite.r.GET("/people", controller.SearchPeopleByTerm)
	suite.r.GET("/people/:id", controller.FindById)
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
type PeopleResponse struct {
	Id         string   `json:"id"`
	Apelido    string   `json:"apelido"`
	Nome       string   `json:"nome"`
	Nascimento string   `json:"nascimento"`
	Stacks     []string `json:"stacks"`
}
