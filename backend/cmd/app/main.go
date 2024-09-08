package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/jibaru/localshops/internal/shared/infrastructure/logger"
	"github.com/jibaru/localshops/internal/shops/application"
	"github.com/jibaru/localshops/internal/shops/infrastructure/handlers"
	"github.com/jibaru/localshops/internal/shops/infrastructure/repositories/orm"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func main() {
	// Log configuration
	logHandler := logger.NewLogHandler(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(slog.New(logHandler))

	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env", "error", err)
		return
	}

	c := DatabaseConfig{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("can not init db", "error", err)
		return
	}

	shopRepo := orm.NewShopRepo(db)
	createShopServ := application.NewCreateServ(shopRepo)
	createShopHandler := handlers.NewCreateHandler(createShopServ)

	// Gin configuration
	gin.SetMode(gin.ReleaseMode)

	// Routes configuration
	r := gin.New()
	r.Use(logger.LogMiddleware)
	r.Use(gin.Recovery())

	r.POST("/shops", createShopHandler.Handle)

	address := "0.0.0.0:8080"

	slog.Info("running app", "address", address)

	if err := r.Run(address); err != nil {
		slog.Error("error running app", "error", err)
	}
}
