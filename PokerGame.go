package main

type PokerGame struct {
    players []Player
    deck
    state GameState
}

func NewPokerGame() *PokerGame {
    game := PokerGame{
        deck: newDeck(),
        players: GeneratePlayers(2),
    }
    game.state = &PreFlopState{game: &game}
    game.state.DealCards()
    game.state.TakeActions()
    return &game
}
