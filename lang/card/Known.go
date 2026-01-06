package card

type DeleteCard interface{
	Delete(pool []int)
}
func (deck*DeckCard) Delete( pool[]int){
	//删除已知牌
   for i:=0; i<len(deck.Cards);{
	    
	  v:=deck.Cards[i]
	  need:=false
	 for _,j:=range pool{
		if v==j{
		  need=true
		  break
		}
	 }
	 if need{
		deck.Cards[i]=deck.Cards[len(deck.Cards)-1]
		deck.Cards=deck.Cards[:len(deck.Cards)-1]
	 }else{
		i++
	 }
   }
  
}
