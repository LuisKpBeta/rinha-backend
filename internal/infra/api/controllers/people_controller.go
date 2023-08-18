package people_controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PeopleController struct {
	createPeople   *people.CreatePeople
	findPeopleById *people.FindPeopleById
}

func CreatePeopleController(create *people.CreatePeople, find *people.FindPeopleById) *PeopleController {
	return &PeopleController{
		createPeople:   create,
		findPeopleById: find,
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

func (p *PeopleController) FindById(c *gin.Context) {
	id, _ := c.Params.Get("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "usuário não encontrado"})
	}
	people, err := p.findPeopleById.Find(userId)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if err == nil && people == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "usuário não encontrado"})
		return
	}

	onlyDate := strings.Split(people.Birthday, "T")[0]
	readPeople := ReadPopleDto{
		Id:       fmt.Sprint(people.Id),
		Nickname: people.Nickname,
		Name:     people.Name,
		Birthday: onlyDate,
		Stacks:   people.GetArrayFromStringStack(),
	}
	c.JSON(http.StatusOK, readPeople)
}
