package main

import (
	gin "github.com/gin-gonic/gin"

	http "net/http"

	strconv "strconv"
)

var products map[int32]Product = map[int32]Product{
	12345: {ProductID: 12345, SKU: "ABC-123-XYZ", Manufacturer: "Boeing", CategoryID: 456, Weight: 41_412_983, SomeOtherID: 789},
	12346: {ProductID: 12346, SKU: "DEF-123-XYZ", Manufacturer: "Airbus", CategoryID: 456, Weight: 37_239_934, SomeOtherID: 790},
}

// handle calls to retrieve an album by its ID
func handleGetProductByID(c *gin.Context) {
	// Validate ID argument
	// TODO: abstract ID validation out from handleGetProductByID and handlePostProductByID
	id64, err := strconv.ParseInt((*c).Param("id"), 10, 32)
	id := int32(id64)
	if err != nil {
		error := Error{
			Error:   "INVALID_ID",
			Message: "The requested ID is not a valid Product ID",
			Details: "Not an integer.",
		}
		(*c).JSON(http.StatusInternalServerError, error)
		return
	}
	if id < 1 {
		error := Error{
			Error:   "INVALID_ID",
			Message: "The requested ID is not a valid Product ID",
			Details: "Less than 1.",
		}
		(*c).JSON(http.StatusInternalServerError, error)
		return
	}

	product, exists := products[id]
	if !exists {
		error := Error{
			Error:   "PRODUCT_NOT_FOUND",
			Message: "There is no product in the database with the given ID.",
		}
		(*c).JSON(http.StatusNotFound, error)
		return
	}
	(*c).JSON(http.StatusOK, product)
}

func handlePostProductByID(c *gin.Context) {
	// Validate ID argument
	id64, err := strconv.ParseInt((*c).Param("id"), 10, 32)
	id := int32(id64)
	if err != nil {
		error := Error{
			Error:   "INVALID_ID",
			Message: "The requested ID is not a valid Product ID",
			Details: "Not an integer.",
		}
		(*c).JSON(http.StatusInternalServerError, error)
		return
	}
	if id < 1 {
		error := Error{
			Error:   "INVALID_ID",
			Message: "The requested ID is not a valid Product ID",
			Details: "Less than 1.",
		}
		(*c).JSON(http.StatusInternalServerError, error)
		return
	}
	// Set product
	var product Product
	// validate productID
	if err := c.ShouldBindJSON(&product); err != nil {
		(*c).JSON(http.StatusBadRequest,
			Error{Error: "INVALID_PRODUCT_DETAILS", Message: err.Error()})
		return
	}
	products[id] = product
	c.Status(http.StatusNoContent)
}

func main() {
	server := gin.Default()
	server.GET("/products/:id", handleGetProductByID)
	server.POST("/products/:id/details", handlePostProductByID)
	// Must run on all interfaces since the container is isolated inside.
	server.Run("0.0.0.0:8080")
}
