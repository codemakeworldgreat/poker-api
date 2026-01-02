package main

import "fmt"
func (deck*Deck)deal(player[]* Player,result *Result,  data*DataCard){
    
   for i:=0;i<result.Person;i++{// 对手发牌
    index:=deck.nextCard(2)
    player[i].Hand=append(player[i].Hand,index...)
    fmt.Println(player[i].Hand)
   }
   fmt.Println(len(deck.Cards))
   if len(data.PublicCard)>0 {
    deck.Hands = append(deck.Hands,data.PublicCard...)
  }
   
    if len(deck.Hands)<7{
     need:=7-len(deck.Hands) 
     outs:=deck.nextCard(need)
     deck.Hands=append(deck.Hands,outs...)
    }

     public:=deck.Hands[2:]
      for i:=0;i<result.Person;i++{
      player[i].Hand=append(player[i].Hand,public...)
      } 
      
      bug(deck.Cards,player)
 }

   
  
   

 


func bitwise(hand[] int) ([]int,[]int){
    //取回数字和花色 
      
      bitnumber:=make([]int,7)
	  bitcolor:=make([]int,7);
	for i:=0;i<len(hand);i++{
       bitnumber[i]=hand[i]>>2;
       bitcolor[i]=hand[i]&3; 
	} 
    
   return bitnumber,bitcolor
	}