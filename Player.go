package main

type PlayerAction int

const (
    Fold PlayerAction = iota
    Check
    Bet
    Call
    Raise
)

type Player struct {
    position int
    cards deck
    stack float64
    currentBet float64
    folded bool
}

func (p *Player) TakeAction(action PlayerAction, betSize float64) (float64){
    switch action {
    case Fold:
        p.folded = true
        return 0.0
    case Call:
        p.currentBet = betSize - p.currentBet
        p.stack = p.stack - p.currentBet
        return p.currentBet
    case Bet:
        p.currentBet = betSize
        p.stack = p.stack - p.currentBet
        return p.currentBet
    case Raise:
        p.currentBet = betSize
        p.stack = p.stack - p.currentBet
        return p.currentBet
    default:
        return 0.0
    }

}

func GeneratePlayers(nPlayers int, buyIn float64) []Player {
    players := []Player{}
    for i := 0; i < nPlayers; i++ {
        players = append(players, Player{
            position: i,
            cards: deck{},
            stack: buyIn,
            currentBet: 0,
            folded: false,
        })
    }
    return players
}
