package people_controller

import (
	"errors"
	"fmt"
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
	var dto CreatePopleDto
	err := c.ShouldBind(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "preencha o body corretamente"})
		return
	}
	err = dto.IsValid()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	created, err := p.createPeople.Create(dto.Nickname, dto.Name, dto.Birthday, dto.Stacks)
	if err != nil {
		if errors.Is(err, people.ErrNickNameAlreadyExists) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, nil)
	location := fmt.Sprint("/pessoas/", created.Id)
	c.Header("Location", location)

}
