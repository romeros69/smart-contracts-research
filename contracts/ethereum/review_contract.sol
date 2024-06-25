// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ReviewContract {
    struct Review {
        string productId;
        address reviewer;
        uint8 rating;
        string comment;
    }

    mapping(string => Review[]) public productReviews;
    mapping(address => mapping(string => bool)) public hasReviewed;

    event ReviewSubmitted(string productId, address indexed reviewer, uint8 rating, string comment);

    function submitReview(string memory productId, uint8 rating, string memory comment) public {
        require(rating >= 1 && rating <= 5, "Rating must be between 1 and 5");
        require(!hasReviewed[msg.sender][productId], "You have already reviewed this product");

        Review memory review = Review({
            productId: productId,
            reviewer: msg.sender,
            rating: rating,
            comment: comment
        });

        productReviews[productId].push(review);
        hasReviewed[msg.sender][productId] = true;

        emit ReviewSubmitted(productId, msg.sender, rating, comment);
    }

    function getProductReviews(string memory productId) public view returns (Review[] memory) {
        return productReviews[productId];
    }

    function getAverageRating(string memory productId) public view returns (uint8) {
        Review[] memory reviews = productReviews[productId];
        uint256 totalRating = 0;
        for (uint256 i = 0; i < reviews.length; i++) {
            totalRating += reviews[i].rating;
        }
        return reviews.length > 0 ? uint8(totalRating / reviews.length) : 0;
    }
}
