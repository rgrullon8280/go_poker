package main

import "fmt"

type GameState interface {
    DealCards()
    TakeActions()
    NextState()
}



func GetPlayerAction(betPlaced bool) int {
    fmt.Println("Actions on you, what do you want to do?")
    fmt.Println("0 - Fold")
    fmt.Println("1 - Call")
    fmt.Println("2 - Check")
    fmt.Println("3 - Bet")
    if betPlaced {
        fmt.Println("4 - Raise")
    }
    var i int
    fmt.Scan(&i)
    return i
}
