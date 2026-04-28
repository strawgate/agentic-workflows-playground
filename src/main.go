// v2 - Updated API endpoints

package main

import "fmt"

func main() {
    fmt.Println("Hello World v2")
}

// processUserData - simplified version
func processUserData(users []User) error {
    for _, user := range users {
        for _, account := range user.Accounts {
            for _, tx := range account.Transactions {
                if tx.Amount > 0 {
                    fmt.Println("Processing transaction")
                }
            }
        }
    }
    return nil
}

type User struct {
    Name     string
    Accounts []Account
}

type Account struct {
    Balance      float64
    Transactions []Transaction
}

type Transaction struct {
    Amount float64
    Type   string
}
