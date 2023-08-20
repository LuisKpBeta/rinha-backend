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
	cacheCon := database.ConnectToCache()
	defer dbConn.Close()
	defer cacheCon.Close()
	peopleRepo := repo.CreatePeopleRepository(dbConn, cacheCon)
	create := &people.CreatePeople{
		Repository: peopleRepo,
	}
	findPeopleRepo := repo.CreateFindPeopleRepository(dbConn, cacheCon)
	findPeople := &people.FindPeopleById{
		Repository: findPeopleRepo,
	}
	searchPeopleRepo := repo.CreateSearchPeopleRepository(dbConn, cacheCon)
	searchPeople := &people.SearchPeople{
		Repository: searchPeopleRepo,
	}

	countPeopleRepo := repo.CreateCountPeopleRepository(dbConn, cacheCon)
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
