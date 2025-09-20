package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (f *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please, enter prices. Confirm every prices with enter")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		} else {
			prices = append(prices, price)
		}
	}
	return prices, nil
}

func (f *CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
