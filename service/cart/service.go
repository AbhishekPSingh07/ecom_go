package cart

import (
	"fmt"

	"github.com/AbhishekPSingh07/ecom_go/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, items := range items {
		if items.Quantity <= 0 {
			return nil, fmt.Errorf("Invalid quantity for the product %d", items.ProductID)
		}

		productIds[i] = items.ProductID
	}
	return productIds,nil
}

func (h *Handler) createOrder(ps []types.Product,item []types.CartItem, userID int) (int ,float64, error){
	
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product	
	}
	items := make([]types.CartCheckoutItems,0)
	//check if the products are actually in stock
	if err := checkIfCartIsInStock(items,productMap); err != nil {
		return 0 ,0 , err
	}
	// calculate the total price
	// reduce the qunatity of products in our db
	// create the order
	// create order items
	return 0,0, nil
}

func checkIfCartIsInStock(cartItems []types.CartCheckoutItems,products map[int]types.Product) error {
	return nil
}
