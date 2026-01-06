package card


type DeckCard struct{
	Cards [] int
	MapCards map[string]int
}

func GetInitial() *DeckCard{ 
	//52张牌的二进制
	deck:=&DeckCard{Cards:make([]int,54),
		MapCards:make(map[string]int),}
     deck.getMap()
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
    return deck
}

   


