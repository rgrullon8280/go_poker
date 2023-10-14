package main

import (
	"fmt"
)

type GameState interface {
    DealCards()
    TakeActions()
    NextState()
}



func GetPlayerAction(betSize float64, playerPosition int, playerStack float64) (PlayerAction, float64) {
    betPlaced := false
    if betSize > 0 {betPlaced = true}
    var i PlayerAction
    fmt.Println("Actions on you Player", playerPosition + 1, ", what do you want to do?")

    validOptions := []string{"0 - Fold"}
    if betPlaced {
        validOptions = append(validOptions, "3 - Call", "4 - Raise")
    } else {
        validOptions = append(validOptions, "1 - Check", "2 - Bet")
    }

    for _, option := range validOptions {
        fmt.Println(option)
    }
    invalidInput := true
    for invalidInput {
        _, err := fmt.Scanln(&i)
        if err != nil {
            fmt.Println("Please enter a valid number")
            ClearInputBuffer()
            continue
        }
        if !betPlaced && (i == Fold || i == Check || i == Bet) {
            invalidInput = false
        } else if betPlaced && (i == Fold || i == Call || i == Raise) {
            invalidInput = false
        } else {
            fmt.Println("Invalid input, please try again.")
        }
    }

    var b float64 = 0.0
    if i == Bet || i == Raise {
        fmt.Println("How much would you like to bet?")
        validBet := false
        for !validBet {
            _, err := fmt.Scanln(&b)
            if err != nil {
                fmt.Println("Please enter a valid number")
                ClearInputBuffer()
                continue
            }
            if b <= 0.0 {
                fmt.Println("Bet size must be positive, pelase try again")
            } else if b > playerStack {
                fmt.Println("Bet size must be less than stack, pelase try again")
            } else {
                validBet = true
            }
        }
    } else if i == Call {b = betSize}
    return i, b
}
