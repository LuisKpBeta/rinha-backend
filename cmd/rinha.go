package main

import (
	"github.com/LuisKpBeta/rinha-backend/internal/infra/api"
	pc "github.com/LuisKpBeta/rinha-backend/internal/infra/api/controllers"
	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
)

func main() {
	create := people.CreatePeople{}
	create.CreatePeople("name", "nickname", "nascimento", make([]string, 0))

	controller := pc.CreatePeopleController(&create)
	r := api.CreateHttpServer()

	r.POST("/people", controller.Create)

	api.StartHttpServer(r)
}
