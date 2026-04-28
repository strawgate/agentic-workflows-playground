// TODO: Implement user authentication
// TODO: Add password reset functionality
// FIXME: Memory leak in connection pool
// TODO: Add unit tests for auth module

package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}

// processUserData has a complex nested loop that could be simplified
func processUserData(users []User) error {
    for i := 0; i < len(users); i++ {
        for j := 0; j < len(users[i].Accounts); j++ {
            for k := 0; k < len(users[i].Accounts[j].Transactions); k++ {
                if users[i].Accounts[j].Transactions[k].Amount > 0 {
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
