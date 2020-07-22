package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/jasongauvin/wikiPattern/models"
)

type config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"3306"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

func HandleRequest(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Poil au message",
		})
	})

	// Listen and serve on 0.0.0.0:8080
	log.Fatal(router.Run(":8080"))
}

func main() {

	HandleRequest()

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	// Database initialization
	models.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
}
