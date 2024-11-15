package boots

import (
	"trinity/config"
	"trinity/database"
	"trinity/modules/campaign"
	"trinity/modules/purchase"
	"trinity/modules/subscription"
	"trinity/modules/voucher"
	"trinity/utils/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

var log logger.Logger

type App struct {
	DB              *mongo.Database
	CampaignHandler *campaign.Handler
	VoucherHandler  *voucher.Handler
	PurchaseHandler *purchase.Handler
}

// InitApp sets up the application dependencies
func InitApp(env *config.Env) (*App, error) {
	db, err := initDB(env)
	if err != nil {
		return nil, err
	}

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

	app := &App{
		DB:              db,
		CampaignHandler: campaignHandler,
		VoucherHandler:  voucherHandler,
		PurchaseHandler: purchaseHandler,
	}

	return app, nil
}

func initDB(env *config.Env) (*mongo.Database, error) {
	log = logger.NewLogger("Initializer")

	dbClient, err := database.NewMongoDB(env)
	if err != nil {
		log.Errorf("failed to connect to MongoDB: %v", err)
		return nil, err
	}

	db := dbClient.Database(env.DBName)
	defer database.CloseMongoDBConnection(dbClient)

	// Setup indexes
	err = database.SetupIndexes(db)
	if err != nil {
		log.Errorf("failed to set up indexes: %v", err)
		return nil, err
	}
	return db, nil
}
