package neo

import (
	"encoding/json"

	"github.com/nspcc-dev/neo-go/pkg/interop"
	"github.com/nspcc-dev/neo-go/pkg/interop/runtime"
	"github.com/nspcc-dev/neo-go/pkg/interop/storage"
)

// Review defines the structure of a review
type ReviewClient struct {
	ProductID string
	Reviewer  interop.Hash160
	Rating    uint8
	Comment   string
}

func Main(operation string, args []interface{}) interface{} {
	switch operation {
	case "submitReview":
		if len(args) != 3 {
			return false
		}
		productId := args[0].(string)
		rating := uint8(args[1].(int))
		comment := args[2].(string)
		return submitReview(productId, rating, comment)
	case "getProductReviews":
		if len(args) != 1 {
			return nil
		}
		productId := args[0].(string)
		return getProductReviews(productId)
	case "getAverageRating":
		if len(args) != 1 {
			return 0
		}
		productId := args[0].(string)
		return getAverageRating(productId)
	default:
		return false
	}
}

func submitReview(productId string, rating uint8, comment string) bool {
	if rating < 1 || rating > 5 {
		return false
	}
	reviewer := runtime.GetCallingScriptHash()
	ctx := storage.GetContext()
	reviewKey := createReviewKey(reviewer, productId)

	if storage.Get(ctx, reviewKey) != nil {
		return false
	}

	review := Review{
		ProductID: productId,
		Reviewer:  string(reviewer),
		Rating:    rating,
		Comment:   comment,
	}

	reviews := getProductReviews(productId)
	reviews = append(reviews, review)

	reviewsBytes, err := json.Marshal(reviews)
	if err != nil {
		return false
	}

	storage.Put(ctx, []byte(productId), reviewsBytes)
	storage.Put(ctx, reviewKey, []byte{1})

	runtime.Notify("ReviewSubmitted", productId, reviewer, rating, comment)
	return true
}

func getProductReviews(productId string) []Review {
	ctx := storage.GetContext()
	reviewsBytes := storage.Get(ctx, []byte(productId))
	if reviewsBytes == nil {
		return []Review{}
	}
	var reviews []Review
	err := json.Unmarshal(reviewsBytes.([]byte), &reviews)
	if err != nil {
		return []Review{}
	}
	return reviews
}

func getAverageRating(productId string) uint8 {
	reviews := getProductReviews(productId)
	if len(reviews) == 0 {
		return 0
	}
	var totalRating uint8
	for _, review := range reviews {
		totalRating += review.Rating
	}
	return totalRating / uint8(len(reviews))
}

func createReviewKey(reviewer interop.Hash160, productId string) []byte {
	return append(reviewer, []byte(productId)...)
}
