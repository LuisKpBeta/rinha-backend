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

	peopleRepo := repo.CreatePeopleRepository(dbConn)
	create := &people.CreatePeople{
		Repository: peopleRepo,
	}

	controller := pc.CreatePeopleController(create)
	r := api.CreateHttpServer()

	r.POST("/people", controller.Create)

	api.StartHttpServer(r)
}
