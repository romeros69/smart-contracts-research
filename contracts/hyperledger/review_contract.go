package hyperledger

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ReviewContract defines the Smart Contract structure
type ReviewContract struct {
	contractapi.Contract
}

// Review defines the structure of a review
type Review struct {
	ProductID string `json:"productId"`
	Reviewer  string `json:"reviewer"`
	Rating    uint8  `json:"rating"`
	Comment   string `json:"comment"`
}

// SubmitReview submits a review for a product
func (rc *ReviewContract) SubmitReview(ctx contractapi.TransactionContextInterface, productId string, rating uint8, comment string) error {
	reviewer, err := rc.getCurrentUser(ctx)
	if err != nil {
		return err
	}

	if rating < 1 || rating > 5 {
		return fmt.Errorf("rating must be between 1 and 5")
	}

	hasReviewed, err := rc.HasReviewed(ctx, reviewer, productId)
	if err != nil {
		return err
	}
	if hasReviewed {
		return fmt.Errorf("you have already reviewed this product")
	}

	review := Review{
		ProductID: productId,
		Reviewer:  reviewer,
		Rating:    rating,
		Comment:   comment,
	}

	reviews, err := rc.GetProductReviews(ctx, productId)
	if err != nil {
		return err
	}

	reviews = append(reviews, review)

	reviewsBytes, err := json.Marshal(reviews)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(productId, reviewsBytes)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(rc.getReviewKey(reviewer, productId), []byte("true"))
	if err != nil {
		return err
	}

	return nil
}

// GetProductReviews returns all reviews for a product
func (rc *ReviewContract) GetProductReviews(ctx contractapi.TransactionContextInterface, productId string) ([]Review, error) {
	reviewsBytes, err := ctx.GetStub().GetState(productId)
	if err != nil {
		return nil, err
	}
	if reviewsBytes == nil {
		return nil, nil
	}

	var reviews []Review
	err = json.Unmarshal(reviewsBytes, &reviews)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

// GetAverageRating returns the average rating for a product
func (rc *ReviewContract) GetAverageRating(ctx contractapi.TransactionContextInterface, productId string) (float32, error) {
	reviews, err := rc.GetProductReviews(ctx, productId)
	if err != nil {
		return 0, err
	}

	if len(reviews) == 0 {
		return 0, nil
	}

	var totalRating uint8
	for _, review := range reviews {
		totalRating += review.Rating
	}

	averageRating := float32(totalRating) / float32(len(reviews))
	return averageRating, nil
}

// HasReviewed checks if a user has reviewed a product
func (rc *ReviewContract) HasReviewed(ctx contractapi.TransactionContextInterface, reviewer, productId string) (bool, error) {
	reviewedBytes, err := ctx.GetStub().GetState(rc.getReviewKey(reviewer, productId))
	if err != nil {
		return false, err
	}
	return reviewedBytes != nil, nil
}

func (rc *ReviewContract) getReviewKey(reviewer, productId string) string {
	return fmt.Sprintf("reviewed_%s_%s", reviewer, productId)
}

func (rc *ReviewContract) getCurrentUser(ctx contractapi.TransactionContextInterface) (string, error) {
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", err
	}
	return clientID, nil
}
