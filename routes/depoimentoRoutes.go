package routes

import (
	"time"

	"github.com/Ulpio/Alura_challange_BackEnd_7/controllers"
	"github.com/Ulpio/Alura_challange_BackEnd_7/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RoutesHandler() {
	database.ConnectDB()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	r.Use(cors.New(config))

	// Endpoints
	RoutesDepoimentos(r)
	RoutesDestinos(r)
	r.Run()
}

func RoutesDepoimentos(r *gin.Engine) {
	r.GET("/depoimentos", controllers.GetDepoimentos)
	r.GET("/depoimentos/:id", controllers.GetDepoimentoPorID)
	r.GET("/depoimentos-home", controllers.GetDepoimentosHome)
	r.POST("/depoimentos", controllers.AdicionarDepoimento)
	r.DELETE("/depoimentos/:id", controllers.DeletarDepoimento)
	r.PUT("/depoimentos/:id", controllers.EditarDepoimento)
}

func RoutesDestinos(r *gin.Engine) {
	r.GET("/destinos", controllers.GetDestinos)
	r.GET("/destinos/:id", controllers.GetDestinoPorID)
	r.GET("/destino", controllers.SearchDestinoByName)
	r.POST("/destinos", controllers.CreateDestino)
	r.DELETE("/destinos/:id", controllers.DeleteDestino)
	r.PATCH("/destinos/:id", controllers.UpdateDestino)
}
