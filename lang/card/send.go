package card 
func (deck*DeckCard)sendCard( need int) []int{
    card:=deck.Cards[:need] 
	deck.Cards=deck.Cards[need:] 
	return card 
}