package controllers

import (
	"net/http"

	"github.com/Ulpio/Alura_challange_BackEnd_7/database"
	"github.com/Ulpio/Alura_challange_BackEnd_7/models"
	"github.com/gin-gonic/gin"
)

func GetDestinos(c *gin.Context) {
	var depoimentos []models.Destinos
	database.DB.Find(&depoimentos)
	c.JSON(http.StatusOK, depoimentos)
}

func GetDestinoPorID(c *gin.Context) {
	id := c.Params.ByName("id")
	var destino models.Destinos
	database.DB.First(&destino, id)
	if destino.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"mensagem": "Destino não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, destino)
}

func SearchDestinoByName(c *gin.Context) {
	nome := c.Query("nome")
	var destinos []models.Destinos
	database.DB.Where("nome = ?", nome).Find(&destinos)
	c.JSON(http.StatusOK, destinos)
}

func CreateDestino(c *gin.Context) {
	var destino models.Destinos
	c.BindJSON(&destino)
	database.DB.Create(&destino)
	c.JSON(http.StatusCreated, destino)
}

func DeleteDestino(c *gin.Context) {
	id := c.Param("id")
	var destino models.Destinos
	database.DB.Delete(&destino, id)
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Destino deletado",
	})
}

func UpdateDestino(c *gin.Context) {
	id := c.Param("id")
	var destino models.Destinos
	database.DB.First(&destino, id)
	c.BindJSON(&destino)
	database.DB.Save(&destino)
	c.JSON(http.StatusOK, gin.H{
		"Mensagem":     "Destino atualizado com sucesso!",
		"Novo destino": destino,
	})
}
