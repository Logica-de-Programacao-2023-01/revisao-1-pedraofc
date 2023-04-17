package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	// 1a=10%
	// >1000=10%
	// <=1000=5%
	// <=500=2%
	// media >1000 = 20%
	// Implemente sua solução aqui
	if len(purchaseHistory) == 0 {
		fmt.Println("Discount:", currentPurchase*0.1)
		return currentPurchase, nil
	} if else purchaseHistory 

}
