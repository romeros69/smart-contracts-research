package neo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"log"

	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	client "github.com/nspcc-dev/neo-go/pkg/rpcclient"
	"github.com/nspcc-dev/neo-go/pkg/util"
)

// Review mirrors the structure used in the smart contract
type Review struct {
	ProductID string
	Reviewer  string
	Rating    uint8
	Comment   string
}

// ReviewContractClient encapsulates RPC client along with contract hash
type ReviewContractClient struct {
	client       *client.Client
	contractHash util.Uint160
}

// NewReviewContractClient creates a new client to interact with the review smart contract
func NewReviewContractClient(rpcURL, contractHash string) (*ReviewContractClient, error) {
	c, err := client.New(context.Background(), rpcURL, client.Options{})
	if err != nil {
		return nil, fmt.Errorf("error creating client: %v", err)
	}
	ch, err := util.Uint160DecodeStringLE(contractHash)
	if err != nil {
		return nil, fmt.Errorf("error decoding contract hash: %v", err)
	}
	return &ReviewContractClient{
		client:       c,
		contractHash: ch,
	}, nil
}

func (r *ReviewContractClient) SubmitReview(productId string, rating uint8, comment string) error {
	_, err := keys.NewPrivateKey()
	if err != nil {
		return fmt.Errorf("error creating account from WIF: %v", err)
	}
	params := []smartcontract.Parameter{
		{
			Type:  smartcontract.StringType,
			Value: productId,
		},
		{
			Type:  smartcontract.IntegerType,
			Value: rating,
		},
		{
			Type:  smartcontract.StringType,
			Value: comment,
		},
	}
	resp, err := r.client.InvokeFunction(r.contractHash, "submitReview", params, nil)
	if err != nil {
		return fmt.Errorf("error invoking submitReview function: %v", err)
	}
	if resp.State != "HALT, BREAK" {
		return fmt.Errorf("execution of submitReview was not successful, state: %v", resp.State)
	}
	fmt.Println("Review submitted successfully")
	return nil
}

func (r *ReviewContractClient) GetProductReviews(productId string) ([]Review, error) {
	params := []smartcontract.Parameter{
		{
			Type:  smartcontract.StringType,
			Value: productId,
		},
	}
	resp, err := r.client.InvokeFunction(r.contractHash, "getProductReviews", params, nil)
	if err != nil {
		return nil, fmt.Errorf("error invoking getProductReviews function: %v", err)
	}
	if len(resp.Stack) == 0 {
		return nil, fmt.Errorf("no data returned")
	}
	data, err := json.Marshal(resp.Stack[0])
	if err != nil {
		return nil, fmt.Errorf("error marshaling response: %v", err)
	}
	var reviews []Review
	if err := json.Unmarshal(data, &reviews); err != nil {
		return nil, fmt.Errorf("error unmarshaling reviews: %v", err)
	}
	return reviews, nil
}

func (r *ReviewContractClient) GetAverageRating(productId string) (uint8, error) {
	params := []smartcontract.Parameter{
		{
			Type:  smartcontract.StringType,
			Value: productId,
		},
	}
	resp, err := r.client.InvokeFunction(r.contractHash, "getAverageRating", params, nil)
	if err != nil {
		return 0, fmt.Errorf("error invoking getAverageRating function: %v", err)
	}
	if len(resp.Stack) == 0 {
		return 0, fmt.Errorf("no data returned")
	}
	averageRating, ok := resp.Stack[0].Value().(uint8)
	if !ok {
		return 0, fmt.Errorf("error asserting rating type")
	}
	return averageRating, nil
}

func main() {
	client, err := NewReviewContractClient("http://localhost:20332", "contract_hash")
	if err != nil {
		log.Fatalf("Error creating contract client: %v", err)
	}

	err = client.SubmitReview("product123", 5, "Great product!")
	if err != nil {
		log.Fatalf("Error submitting review: %v", err)
	}

	reviews, err := client.GetProductReviews("product123")
	if err != nil {
		log.Fatalf("Error retrieving reviews: %v", err)
	}
	fmt.Printf("Reviews: %+v\n", reviews)

	averageRating, err := client.GetAverageRating("product123")
	if err != nil {
		log.Fatalf("Error getting average rating: %v", err)
	}
	fmt.Printf("Average Rating: %d\n", averageRating)
}
