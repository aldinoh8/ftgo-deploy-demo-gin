package controller

import (
	"aldinoh8/example-gin/entity"
	"aldinoh8/example-gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Main struct {
	Repository repository.Main
}

func New(r repository.Main) Main {
	return Main{Repository: r}
}

func (c Main) Create(ctx *gin.Context) {
	body := entity.Todo{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create todo",
			"detail":  err.Error(),
		})
		return
	}

	todo, err := c.Repository.Create(body.Task, body.Description)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create todo",
			"detail":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success create todo",
		"todo":    todo,
	})
}

func (c Main) FindAll(ctx *gin.Context) {
	todos, err := c.Repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to retreive todos",
			"detail":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}
