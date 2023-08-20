package main

import (
	"github.com/LuisKpBeta/rinha-backend/internal/infra/api"
	pc "github.com/LuisKpBeta/rinha-backend/internal/infra/api/controllers"
	"github.com/LuisKpBeta/rinha-backend/internal/infra/database"
	repo "github.com/LuisKpBeta/rinha-backend/internal/infra/database/people"
	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
)

func main() {
	dbConn := database.ConnectToDatabase()
	defer dbConn.Close()
	peopleRepo := repo.CreatePeopleRepository(dbConn)
	create := &people.CreatePeople{
		Repository: peopleRepo,
	}
	findPeopleRepo := repo.CreateFindPeopleRepository(dbConn)
	findPeople := &people.FindPeopleById{
		Repository: findPeopleRepo,
	}
	searchPeopleRepo := repo.CreateSearchPeopleRepository(dbConn)
	searchPeople := &people.SearchPeople{
		Repository: searchPeopleRepo,
	}

	countPeopleRepo := repo.CreateCountPeopleRepository(dbConn)
	countPeople := &people.CountPeople{
		Repository: countPeopleRepo,
	}

	controller := pc.CreatePeopleController(create, findPeople, searchPeople, countPeople)
	r := api.CreateHttpServer()

	r.POST("/pessoas", controller.Create)
	r.GET("/pessoas", controller.SearchPeopleByTerm)
	r.GET("/pessoas/:id", controller.FindById)
	r.GET("/contagem-pessoas", controller.Count)

	api.StartHttpServer(r)
}
