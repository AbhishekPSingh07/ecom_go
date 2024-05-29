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

func (h *Handler) createOrder(ps []types.Product,items []types.CartItem, userID int) (int ,float64, error){
	
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product	
	}

	//check if the products are actually in stock
	if err := checkIfCartIsInStock(items,productMap); err != nil {
		return 0 ,0 , nil
	}
	// calculate the total price
	totalPrice := calculateTotalPrice(items,productMap)
	// reduce the qunatity of products in our db
	for _,item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}
	// create the order
	orderId, err := h.store.CreateOrder(types.Order{
		UserID: userID,
		Total: totalPrice,
		Status: "pending",
		Address: "some Address",
	})
	if err != nil {
		return 0 ,0,nil
	}
	// create order items
	for _ ,item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID : orderId,
			ProductID : item.ProductID,
			Quantity : item.Quantity,
			Price: productMap[item.ProductID].Price,
		})
	}
	return orderId,totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem,products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _,item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %s is not available in the quantity requested",product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartItem,products map[int]types.Product) (float64) {
    var total float64
	for _, item := range cartItems {
		product := products[item.ProductID]
		total += float64(item.Quantity)*product.Price
	}

	return total
}
