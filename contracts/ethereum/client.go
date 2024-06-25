package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ReviewContractClient struct {
	client          *ethclient.Client
	contract        *Ethereum
	auth            *bind.TransactOpts
	contractAddress common.Address
}

func NewReviewContractClient(url, contractAddress string, privateKey *ecdsa.PrivateKey) (*ReviewContractClient, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)
	contract, err := NewEthereum(address, client)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // Укажите chainID вашей сети
	if err != nil {
		return nil, err
	}

	return &ReviewContractClient{
		client:          client,
		contract:        contract,
		auth:            auth,
		contractAddress: address,
	}, nil
}

func (c *ReviewContractClient) SubmitReview(productId string, rating uint8, comment string) error {
	tx, err := c.contract.SubmitReview(c.auth, productId, rating, comment)
	if err != nil {
		return err
	}

	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
	return nil
}

func (c *ReviewContractClient) GetProductReviews(productId string) ([]ReviewContractReview, error) {
	return c.contract.GetProductReviews(&bind.CallOpts{}, productId)
}

func (c *ReviewContractClient) GetAverageRating(productId string) (uint8, error) {
	return c.contract.GetAverageRating(&bind.CallOpts{}, productId)
}

func main() {
	privateKey, err := crypto.HexToECDSA("private-key")
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	client, err := NewReviewContractClient("http://localhost:8545", "address", privateKey)
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
