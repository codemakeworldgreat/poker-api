package main
import "math/rand"

func (deck *Deck) initial() { //初始化
	colors:=[]int{0,1,2,3};
	numbers:=make([]int,13);
	for i:=2;i<15;i++{
		numbers[i-2]=i;
	}
	index:=0;
	for i:=0;i<len(numbers);i++{
		for j:=0;j<len(colors);j++{
			deck.Cards[index]=(numbers[i]<<2)|colors[j];
			index++;
		}
	}
   
}

func (deck *Deck) know(data *DataCard, player[]*Player, cards ...int) {
	   //去掉已知牌
	
	 
	for f, t := range cards {
          deck.Hands[f]=t 
		for i, v := range deck.Cards {
			if v == t {
				deck.Cards[i] = deck.Cards[len(deck.Cards)-1]
				deck.Cards = deck.Cards[:len(deck.Cards)-1]	
				break
			}
		}
	}
      
	if len(data.PublicCard) !=0 { //如果公牌已知 删除并给每个玩家加上公牌
	    
		
		for _, v := range data.PublicCard {
			for i, j := range deck.Cards {
				if v == j {
			        deck.Cards[i]=deck.Cards[len(deck.Cards)-1]
					deck.Cards=deck.Cards[:len(deck.Cards)-1]
					  
                           break
                        }
				}
			}
		}
	}



func (deck *Deck) shuffle() {//洗牌
    cards:=deck.Cards
    rand.Shuffle(len(cards),func(i,j int){
		cards[i],cards[j]=cards[j],cards[i]
	})	
    deck.Cards=cards
    
}
func (deck*Deck) nextCard(index int)[]int{
	
	needcards:=deck.Cards[:index]
    deck.Cards=deck.Cards[index:]
	return  needcards
}
   


