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
    currentBet := 0.0
    for i := range p.game.players {
        player := &p.game.players[i]
        playerAction, betSize := GetPlayerAction(currentBet, i, player.stack)
        currentBet = betSize
        player.TakeAction(playerAction, betSize)
        fmt.Println("Current bet is:", betSize)
        fmt.Println("Player:", *player)
    }
}

func (p PreFlopState) NextState() {
    fmt.Println("Next State")
} 
