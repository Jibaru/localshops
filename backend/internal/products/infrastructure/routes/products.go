package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	shops_application "github.com/jibaru/localshops/internal/shops/application"
	shops_orm "github.com/jibaru/localshops/internal/shops/infrastructure/repositories/orm"

	"github.com/jibaru/localshops/internal/products/application"
	"github.com/jibaru/localshops/internal/products/domain"
	"github.com/jibaru/localshops/internal/products/infrastructure/handlers"
	remoteservices "github.com/jibaru/localshops/internal/products/infrastructure/remote_services"
	"github.com/jibaru/localshops/internal/products/infrastructure/repositories/orm"
)

func Set(r *gin.Engine, db *gorm.DB) {
	shopRepo := shops_orm.NewShopRepo(db)
	getShopServ := shops_application.NewGetServ(shopRepo)

	productRepo := orm.NewProductRepo(db)
	externalShopServ := remoteservices.NewShopServ(getShopServ)
	createProductDomainServ := domain.NewProductServ(externalShopServ)
	createProductServ := application.NewCreateServ(productRepo, createProductDomainServ)

	createProductHandler := handlers.NewCreateHandler(createProductServ)

	r.POST("/products", createProductHandler.Handle)
}
