package card

func (deck*DeckCard) delete( pool[]int){
	//删除已知牌
   for i,v:=range deck.Cards{
   for _,j:=range pool{
	if v==j{
		deck.Cards[i]=deck.Cards[len(deck.Cards)-1]
        deck.Cards=deck.Cards[:len(deck.Cards)-1]
		break
	}
   }
   }
  
}
