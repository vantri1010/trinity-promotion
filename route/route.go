package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"trinity/bootstrap"
	_ "trinity/docs"
	"trinity/modules/campaign"
	"trinity/modules/purchase"
	"trinity/modules/subscription"
	"trinity/modules/voucher"
	"trinity/mongo"
)

func SetupRouter(env *bootstrap.Env, db mongo.Database, router *gin.Engine) {
	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	api := router.Group("/")

	// Repositories
	campaignRepo := campaign.NewRepository(db)
	voucherRepo := voucher.NewRepository(db)
	subscriptionRepo := subscription.NewRepository(db)
	purchaseRepo := purchase.NewRepository(db)

	// Services
	campaignService := campaign.NewService(campaignRepo, voucherRepo)
	voucherService := voucher.NewService(voucherRepo)
	purchaseService := purchase.NewService(purchaseRepo, voucherRepo, subscriptionRepo)

	// Handlers
	campaignHandler := campaign.NewHandler(campaignService)
	voucherHandler := voucher.NewHandler(voucherService)
	purchaseHandler := purchase.NewHandler(purchaseService)

	// Campaign routes
	campaignRoutes := api.Group("/campaigns")
	campaignHandler.RegisterRoutes(campaignRoutes)

	// Voucher routes
	voucherRoutes := api.Group("/vouchers")
	voucherHandler.RegisterRoutes(voucherRoutes)

	// Purchase routes
	purchaseRoutes := api.Group("/purchases")
	purchaseHandler.RegisterRoutes(purchaseRoutes)

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
