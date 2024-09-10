package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/jibaru/localshops/internal/shops/application"
	"github.com/jibaru/localshops/internal/shops/infrastructure/handlers"
	"github.com/jibaru/localshops/internal/shops/infrastructure/repositories/orm"
)

func Set(r *gin.Engine, db *gorm.DB) {
	shopRepo := orm.NewShopRepo(db)
	createShopServ := application.NewCreateServ(shopRepo)
	getAllShopsServ := application.NewGetAllServ(shopRepo)
	createShopHandler := handlers.NewCreateHandler(createShopServ)
	getAllShopsHandler := handlers.NewGetAllHandler(getAllShopsServ)

	r.POST("/shops", createShopHandler.Handle)
	r.GET("/shops", getAllShopsHandler.Handle)
}
