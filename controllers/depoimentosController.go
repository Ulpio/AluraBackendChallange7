package controllers

import (
	"net/http"

	"github.com/Ulpio/Alura_challange_BackEnd_7/database"
	"github.com/Ulpio/Alura_challange_BackEnd_7/models"
	"github.com/gin-gonic/gin"
)

func GetDepoimentos(c *gin.Context) {
	var depoimentos []models.Depoimentos
	database.DB.Find(&depoimentos)
	c.JSON(http.StatusOK, depoimentos)
}

func GetDepoimentoPorID(c *gin.Context) {
	id := c.Params.ByName("id")
	var depoimento models.Depoimentos
	database.DB.First(&depoimento, id)
	c.JSON(http.StatusOK, depoimento)
}

func GetDepoimentosHome(c *gin.Context) {
	// Mostre 3 depoimentos aleatorios
	// E que nao se repitam
	var depoimentos []models.Depoimentos
	database.DB.Order("RANDOM()").Limit(3).Find(&depoimentos)
	c.JSON(http.StatusOK, depoimentos)
}

func AdicionarDepoimento(c *gin.Context) {
	var depoimento models.Depoimentos
	c.BindJSON(&depoimento)
	database.DB.Create(&depoimento)
	c.JSON(http.StatusCreated, depoimento)
}

func DeletarDepoimento(c *gin.Context) {
	id := c.Params.ByName("id")
	var depoimento models.Depoimentos
	database.DB.Delete(&depoimento, id)
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Depoimento de id" + id + "deletado com sucesso",
	})
}

func EditarDepoimento(c *gin.Context) {
	id := c.Params.ByName("id")
	var depoimento models.Depoimentos
	database.DB.First(&depoimento, id)
	c.BindJSON(&depoimento)
	database.DB.Save(&depoimento)
	c.JSON(http.StatusOK, depoimento)
}
