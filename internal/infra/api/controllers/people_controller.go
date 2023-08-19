package people_controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/LuisKpBeta/rinha-backend/internal/services/people"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PeopleController struct {
	createPeople   *people.CreatePeople
	findPeopleById *people.FindPeopleById
	searchPeople   *people.SearchPeople
}

func CreatePeopleController(create *people.CreatePeople, find *people.FindPeopleById, search *people.SearchPeople) *PeopleController {
	return &PeopleController{
		createPeople:   create,
		findPeopleById: find,
		searchPeople:   search,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if err == nil && people == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "usuário não encontrado"})
		return
	}

	readPeople := ReadPopleDto{
		Id:       fmt.Sprint(people.Id),
		Nickname: people.Nickname,
		Name:     people.Name,
		Birthday: people.GetBirthdayFormated(),
		Stacks:   people.GetArrayFromStringStack(),
	}
	c.JSON(http.StatusOK, readPeople)
}

func (p *PeopleController) SearchPeopleByTerm(c *gin.Context) {
	term := c.Query("t")
	if term == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "o query param 't' é obrigátorio para busca"})
	}
	pList, err := p.searchPeople.SearchByTerm(term)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	peopleList := make([]ReadPopleDto, len(pList))
	if len(pList) == 0 {
		c.JSON(http.StatusOK, peopleList)
		return
	}
	for i, pItem := range pList {
		readPeople := ReadPopleDto{
			Id:       fmt.Sprint(pItem.Id),
			Nickname: pItem.Nickname,
			Name:     pItem.Name,
			Birthday: pItem.GetBirthdayFormated(),
			Stacks:   pItem.GetArrayFromStringStack(),
		}
		peopleList[i] = readPeople
	}
	c.JSON(http.StatusOK, peopleList)
}
