package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ekokurniadi/pos-golang/entity"
	apiHandler "github.com/ekokurniadi/pos-golang/handler"
	"github.com/ekokurniadi/pos-golang/repository"
	"github.com/ekokurniadi/pos-golang/service"
	"github.com/ekokurniadi/pos-golang/web/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Gagal mendapatkan environment file")
	}

	host := os.Getenv("DB_HOST")
	userHost := os.Getenv("DB_USER")
	userPass := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")
	databasePort := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + userHost + " password=" + userPass + " dbname=" + databaseName + " port=" + databasePort + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&entity.Item{})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userWebHandler := handler.NewSessionHandler(userService)
	itemRepository := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepository)

	itemHandler := apiHandler.NewItemHandler(itemService)

	transaksiWebHandler := handler.NewTransaksiHandler()
	noRouteHandler := handler.NewNotRouteHandler()

	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("web/templates/**/*")
	router.Static("css", "./web/assets/css")
	router.Static("js", "./web/assets/js")
	router.Static("bootstrap", "./web/assets/bootstrap")
	router.Static("webfonts", "./web/assets/webfonts")
	router.Static("images", "./images")
	api := router.Group("/api/v1")

	api.GET("/items", itemHandler.GetItems)
	api.POST("/items", itemHandler.CreateItem)
	api.POST("/items/search", itemHandler.SearchItem)
	api.GET("/items/:id", itemHandler.GetItem)

	router.GET("/dashboard", userWebHandler.Dashboard)
	router.GET("/transaksi", transaksiWebHandler.Index)
	router.NoRoute(noRouteHandler.Index)
	router.Run(":8000")

	fmt.Println("Database Connected")
}
