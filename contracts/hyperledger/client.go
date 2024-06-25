package hyperledger

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

// ReviewContractClient wraps the SDK's functionality for easier use.
type ReviewContractClient struct {
	gateway  *gateway.Gateway
	network  *gateway.Network
	contract *gateway.Contract
}

// NewReviewContractClient creates a new instance of ReviewContractClient.
func NewReviewContractClient(configPath, walletPath, user, channelName, contractName string) (*ReviewContractClient, error) {
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	if !wallet.Exists(user) {
		return nil, fmt.Errorf("user not found in wallet: %s", user)
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(configPath)),
		gateway.WithIdentity(wallet, user),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gateway: %w", err)
	}

	network, err := gw.GetNetwork(channelName)
	if err != nil {
		return nil, fmt.Errorf("failed to get network: %w", err)
	}

	contract := network.GetContract(contractName)

	return &ReviewContractClient{
		gateway:  gw,
		network:  network,
		contract: contract,
	}, nil
}

// SubmitReview submits a review for a product.
func (c *ReviewContractClient) SubmitReview(productId string, rating uint8, comment string) error {
	_, err := c.contract.SubmitTransaction("SubmitReview", productId, fmt.Sprintf("%d", rating), comment)
	return err
}

// GetProductReviews retrieves the reviews for a product.
func (c *ReviewContractClient) GetProductReviews(productId string) ([]Review, error) {
	result, err := c.contract.EvaluateTransaction("GetProductReviews", productId)
	if err != nil {
		return nil, err
	}

	var reviews []Review
	err = json.Unmarshal(result, &reviews)
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetAverageRating retrieves the average rating for a product.
func (c *ReviewContractClient) GetAverageRating(productId string) (float32, error) {
	result, err := c.contract.EvaluateTransaction("GetAverageRating", productId)
	if err != nil {
		return 0, err
	}

	var averageRating float32
	err = json.Unmarshal(result, &averageRating)
	if err != nil {
		return 0, err
	}
	return averageRating, nil
}

func main() {
	client, err := NewReviewContractClient("config.yaml", "wallet", "user1", "mychannel", "review-contract")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	err = client.SubmitReview("product123", 5, "Great product!")
	if err != nil {
		log.Fatalf("Failed to submit review: %v", err)
	}

	reviews, err := client.GetProductReviews("product123")
	if err != nil {
		log.Fatalf("Failed to get reviews: %v", err)
	}

	fmt.Println("Reviews:", reviews)

	avgRating, err := client.GetAverageRating("product123")
	if err != nil {
		log.Fatalf("Failed to get average rating: %v", err)
	}

	fmt.Println("Average Rating:", avgRating)
}
