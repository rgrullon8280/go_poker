package main

type PokerGame struct {
    players []Player
    deck
    state GameState
    pot float64
    smallBlind float64
    bigBlind float64
    buttonPosition int
}

func NewPokerGame(smallBlind float64, bigBlind float64, buyIn float64, numberOfPlayers int) *PokerGame {
    game := PokerGame{
        deck: newDeck(),
        players: GeneratePlayers(numberOfPlayers, buyIn),
        smallBlind: smallBlind,
        bigBlind: bigBlind,
        pot: smallBlind + bigBlind,
        buttonPosition: 0,
    }
    game.state = &PreFlopState{game: &game}
    game.state.DealCards()
    game.state.TakeActions()
    return &game
}

func (pg *PokerGame) FirstToActPreFlop() int {
    var firstToAct int
    if len(pg.players) == 2 {
        firstToAct = pg.buttonPosition
    } else {
        firstToAct = (pg.buttonPosition + 1) % len(pg.players)
    }
    return firstToAct
}

func (pg *PokerGame) RotatePositions() {
    if len(pg.players) == 2 {
        pg.buttonPosition = 1 - pg.buttonPosition // toggle between 0 and 1
    } else {
        pg.buttonPosition = (pg.buttonPosition + 1) % len(pg.players)
    }

    for i, _ := range pg.players {
        pg.players[i].position = (pg.buttonPosition + i) % len(pg.players)
    }
}

