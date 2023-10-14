package main

type PokerGame struct {
    players []Player
    deck
    state GameState
    pot float64
    smallBlind float64
    bigBlind float64
}

func NewPokerGame(smallBlind float64, bigBlind float64, buyIn float64, numberOfPlayers int) *PokerGame {
    game := PokerGame{
        deck: newDeck(),
        players: GeneratePlayers(numberOfPlayers, buyIn),
        smallBlind: smallBlind,
        bigBlind: bigBlind,
        pot: smallBlind + bigBlind,
    }
    game.state = &PreFlopState{game: &game}
    game.state.DealCards()
    game.state.TakeActions()
    return &game
}
