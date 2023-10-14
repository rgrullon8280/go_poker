package main

type PlayerAction int

const (
    Fold PlayerAction = iota
    Check
    Call
    Bet
    Raise
)

type Player struct {
    position int
    cards deck
    stack float32
    currentBet float32
    folded bool
}

func (p *Player) TakeAction() {

}

func GeneratePlayers(nPlayers int) []Player {
    players := []Player{}
    for i := 0; i < nPlayers; i++ {
        players = append(players, Player{
            position: i,
            cards: deck{},
            stack: 0,
            currentBet: 0,
            folded: false,
        })
    }
    return players
}
