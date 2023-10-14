package main

type PostFlopState struct {
    game *PokerGame
    cardRound CardRound
}

func (pf *PostFlopState) DealCards() {

}
func (pf *PostFlopState) TakeActions() {

}
func (pf *PostFlopState) NextState() GameState {
    var nextState GameState
    if pf.cardRound < 2 {
        nextState = &PostFlopState{
            game: pf.game,
            cardRound: pf.cardRound + 1,
        }
    } else {
        nextState = &PreFlopState{game: pf.game}
    }
    return nextState
}

