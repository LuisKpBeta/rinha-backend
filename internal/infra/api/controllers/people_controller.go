package people_controller

import (
	"net/http"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/gin-gonic/gin"
)

type PeopleController struct {
	createPeople *people.CreatePeople
}

func CreatePeopleController(create *people.CreatePeople) *PeopleController {
	return &PeopleController{
		createPeople: create,
	}
}

func (p *PeopleController) Create(c *gin.Context) {
	dto := &CreatePopleDto{}
	err := c.Bind(&dto)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	p.createPeople.CreatePeople(dto.Name, dto.Nickname, dto.Birthday, dto.Stacks)

}
