package main
import (
    "fmt"
)

type PreFlopState struct {
    game *PokerGame
}

func (p *PreFlopState) DealCards() {
    for i := range p.game.players {
        p.game.players[i].cards = p.game.deck.deal(2)
    }
}

func (p *PreFlopState) TakeActions() {
    betPlaced := false
    for i := range p.game.players {
        player := &p.game.players[i]
        playerAction := GetPlayerAction(betPlaced)
        fmt.Println(player)
        fmt.Println(playerAction)
    }
}

func (p PreFlopState) NextState() {
    fmt.Println("Next State")
} 
