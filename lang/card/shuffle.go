package card
import "math/rand/v2"

func (deck*DeckCard)Shuffle(){
   rand.Shuffle(len(deck.Cards),func(i,j int){
     deck.Cards[i],deck.Cards[j]=deck.Cards[j],deck.Cards[i]
   })
}