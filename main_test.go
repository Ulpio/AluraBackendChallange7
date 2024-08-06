package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ulpio/Alura_challange_BackEnd_7/controllers"
	"github.com/Ulpio/Alura_challange_BackEnd_7/database"
	"github.com/Ulpio/Alura_challange_BackEnd_7/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	ID int
)

func SetupRoutesTest() *gin.Engine {
	database.ConnectDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func CriaDestinoMock() {
	destino := models.Destinos{
		Nome:  "Teste",
		Preco: 100,
		Foto:  "https://www.google.com",
	}
	database.DB.Create(&destino)
	ID = int(destino.ID)
}

func DeletaDestinoMock() {
	var destino models.Destinos
	database.DB.Delete(&destino, ID)
}

func TestGetDestinos(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/destinos", controllers.GetDestinos)
	req, _ := http.NewRequest("GET", "/destinos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetDestinosPorID(t *testing.T) {
	database.ConnectDB()
	CriaDestinoMock()
	defer DeletaDestinoMock()
	r := SetupRoutesTest()
	r.GET("/destinos/:id", controllers.GetDestinoPorID)
	req, _ := http.NewRequest("GET", "/destinos/"+fmt.Sprint(ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestSearchDestinoByName(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/destino", controllers.SearchDestinoByName)
	req, _ := http.NewRequest("GET", "/destino?nome=Teste", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteDestino(t *testing.T) {
	database.ConnectDB()
	CriaDestinoMock()
	r := SetupRoutesTest()
	r.DELETE("/destinos/:id", controllers.DeleteDestino)
	req, _ := http.NewRequest("DELETE", "/destinos/"+fmt.Sprint(ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateDestino(t *testing.T) {
	database.ConnectDB()
	CriaDestinoMock()
	defer DeletaDestinoMock()
	r := SetupRoutesTest()
	r.PATCH("/destinos/:id", controllers.UpdateDestino)
	destino := models.Destinos{Nome: "Nome de Teste", Preco: 175, Foto: "https://www.google.com/foto"}
	valorJson, _ := json.Marshal(destino)
	req, _ := http.NewRequest("PATCH", "/destinos/"+fmt.Sprint(ID), bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateDestino(t *testing.T) {
	database.ConnectDB()
	defer DeletaDestinoMock()
	r := SetupRoutesTest()
	r.POST("/destinos", controllers.CreateDestino)
	destino := models.Destinos{Nome: "Nome de Teste", Preco: 175, Foto: "https://www.google.com/foto", Foto2: "https://www.google.com/foto2", Meta: "Meta", TextoDescritivo: "Texto Descritivo"}
	valorJson, _ := json.Marshal(destino)
	req, _ := http.NewRequest("POST", "/destinos", bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	var destinoCriado models.Destinos
	json.NewDecoder(resp.Body).Decode(&destinoCriado)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Equal(t, destino.Nome, destinoCriado.Nome)
	assert.Equal(t, destino.Preco, destinoCriado.Preco)
	assert.Equal(t, destino.Foto, destinoCriado.Foto)
	database.DB.Delete(&destinoCriado)
}

func CreateDepoimentoMock() {
	depoimento := models.Depoimentos{
		Autor: "Teste",
		Texto: "Testando",
		Foto:  "https://www.google.com",
	}
	database.DB.Create(&depoimento)
	ID = int(depoimento.ID)
}

func DeleteDepoimentoMock() {
	var depoimento models.Depoimentos
	database.DB.Delete(&depoimento, ID)
}
func TestGetDepoimentos(t *testing.T) {
	database.ConnectDB()
	r := SetupRoutesTest()
	r.GET("/depoimentos", controllers.GetDepoimentos)
	req, _ := http.NewRequest("GET", "/depoimentos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetDepoimentoPorID(t *testing.T) {
	database.ConnectDB()
	CreateDepoimentoMock()
	defer DeleteDepoimentoMock()
	r := SetupRoutesTest()
	r.GET("/depoimentos/:id", controllers.GetDepoimentoPorID)
	req, _ := http.NewRequest("GET", "/depoimentos/"+fmt.Sprint(ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetDepoimentosHome(t *testing.T) {
	database.ConnectDB()
	r := SetupRoutesTest()
	r.GET("/depoimentos-home", controllers.GetDepoimentosHome)
	req, _ := http.NewRequest("GET", "/depoimentos-home", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteDepoimento(t *testing.T) {
	database.ConnectDB()
	CreateDepoimentoMock()
	r := SetupRoutesTest()
	r.DELETE("/depoimentos/:id", controllers.DeletarDepoimento)
	req, _ := http.NewRequest("DELETE", "/depoimentos/"+fmt.Sprint(ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestEditDepoimento(t *testing.T) {
	database.ConnectDB()
	CreateDepoimentoMock()
	defer DeleteDepoimentoMock()
	r := SetupRoutesTest()
	r.PUT("/depoimentos/:id", controllers.EditarDepoimento)
	depoimento := models.Depoimentos{Autor: "Nome de Teste", Texto: "Testando", Foto: "https://www.google.com/foto"}
	valorJson, _ := json.Marshal(depoimento)
	req, _ := http.NewRequest("PUT", "/depoimentos/"+fmt.Sprint(ID), bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestAdicionarDepoimento(t *testing.T) {
	database.ConnectDB()
	defer DeleteDepoimentoMock()
	r := SetupRoutesTest()
	r.POST("/depoimentos", controllers.AdicionarDepoimento)
	depoimento := models.Depoimentos{Autor: "Nome de Teste", Texto: "Testando", Foto: "https://www.google.com/foto"}
	valorJson, _ := json.Marshal(depoimento)
	req, _ := http.NewRequest("POST", "/depoimentos", bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)
}
