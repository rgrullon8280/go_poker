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
    currentPlayer := p.game.FirstToActPreFlop()
    currentBet := 0.0
    for i := range p.game.players {
        player := &p.game.players[currentPlayer]
        playerAction, betSize := GetPlayerAction(currentBet, i, player.stack)
        currentPlayer = (currentPlayer + 1) % len(p.game.players)
        currentBet = betSize
        player.TakeAction(playerAction, betSize)
        fmt.Println("Current bet is:", betSize)
        fmt.Println("Player:", *player)
    }
}

func (p *PreFlopState) NextState() GameState {
    return &PostFlopState{
                        game: p.game,
                        cardRound: 0,
    }

} 
