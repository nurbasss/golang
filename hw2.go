package main

import (
	"fmt"
	"log"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println(Price.String(25945))
	fmt.Println(Prices)
	RegisterItem(Prices,"vodka",25)
	cart := new(Cart)
	cart.AddItem("choclate")
	fmt.Println(cart.hasMilk())
	fmt.Println(cart.HasItem("doner"))

	cart.AddItem("eggs")
	cart.Checkout()
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	
	return fmt.Sprintf("$%.2f", float32(p)/100)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	if _, exist := prices[item]; exist {
		log.Printf("It's exists, price updated to %s\n",
			 price.String())
	} else {
		fmt.Printf("Added new item. %s: %s\n", item, price.String())
	}
	prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	for _, item := range c.Items {
		if item == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, i := range c.Items {
		if i == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	if price, ok := Prices[item]; ok {
		c.Items = append(c.Items, item)
		c.TotalPrice += price
	} else {
		log.Printf("There is no price for %v\n", item)
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Println("Your cart:")
	for _, i := range c.Items {
		fmt.Println(i, Prices[i])
	}

	fmt.Printf("Total price: %v\n", c.TotalPrice)

	c.Items = c.Items[:0]
	c.TotalPrice = 0
}