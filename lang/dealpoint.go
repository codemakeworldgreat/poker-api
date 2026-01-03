package main

type Myhand struct{
	Color []int
	Number []int
	
}
  
func dealpoint(deck *Deck,player[]*Player)(float32) {
     win:=float32(0)
	 number,color:=bitwise(deck.Hands)
	 myhand:=Myhand{Color:color,Number:number}
	 mypoint:=dealnumber(myhand.Number,myhand.Color)//改到这里
     
	 
	 for i:=0;i<len(player);i++{
          
         playnumber,playcolor:= bitwise(player[i].Hand)
		 
		 player[i].Color=playcolor
		 player[i].Number=playnumber
       player[i].Points=dealnumber(player[i].Number,player[i].Color)
	 }
	  
     win +=pk(&mypoint,player)
	 
	return win
}
