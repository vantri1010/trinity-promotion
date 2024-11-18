package campaign

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"trinity/modules/model"

	"trinity/mongo"
)

// Repository defines campaign data access methods
type Repository interface {
	CreateCampaign(campaign *model.Campaign) (string, error)
	GetCampaignByID(id string) (*model.Campaign, error)
	IncrementUsedUsers(id string, count int) error
	ListCampaigns() ([]model.Campaign, error)
}

// repository implements Repository interface
type repository struct {
	collection mongo.Collection
}

// NewRepository creates a new Campaign repository
func NewRepository(db mongo.Database) Repository {
	return &repository{
		collection: db.Collection("campaigns"),
	}
}

// CreateCampaign inserts a new campaign into the database and returns its ID
func (r *repository) CreateCampaign(campaign *model.Campaign) (string, error) {
	result, err := r.collection.InsertOne(context.Background(), campaign)
	if err != nil {
		return "", err
	}

	oid, ok := result.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert to ObjectID")
	}

	campaign.Id = oid.Hex()
	return campaign.Id, nil
}

// GetCampaignByID retrieves a campaign by its ID
func (r *repository) GetCampaignByID(campaignID string) (*model.Campaign, error) {
	objID, err := primitive.ObjectIDFromHex(campaignID)
	if err != nil {
		return nil, fmt.Errorf("invalid campaign ID format")
	}

	var campaign model.Campaign
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&campaign)

	return &campaign, err
}

// IncrementUsedUsers increments the used_users field of a campaign by a specified count
func (r *repository) IncrementUsedUsers(id string, count int) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid campaign ID: %w", err)
	}

	update := bson.M{
		"$inc": bson.M{
			"used_users": count,
		},
	}

	result, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return fmt.Errorf("failed to increment used users: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no campaign found with ID %s", id)
	}

	return nil
}

// ListCampaigns retrieves all campaigns from the database
func (r *repository) ListCampaigns() ([]model.Campaign, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var campaigns []model.Campaign

	err = cursor.All(context.Background(), &campaigns)

	if err != nil {
		return []model.Campaign{}, err
	}

	return campaigns, nil
}
