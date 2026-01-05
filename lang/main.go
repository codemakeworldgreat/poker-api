package main

import (
	"fmt"
	"poker/card"
    "poker/deal"
)
type DataCard struct{
	Card1 int
	Card2 int
	PublicCard []int
	Person int 
    
}

 
func main(){
  
 // requestcard();//启动接收数据从laravel
       
     
      p:=&card.Player{Index: make([]card.Detail,8),}
	  tb:=card.Table{}
      d:=card.GetInitial()
	 public:=[]string{"AC","6H","5C","10H","QC"}
      data:=d.SetMap(public)
	  hand:=[]string{"8H","8S"}
      hand3:=[]string{"JH","JC"}
      hands:=d.SetMap(hand)
	  hands3:=d.SetMap(hand3)
	  fmt.Println(data,hands)
	  d.Delete(hands)
	  d.Delete(data)
	  d.Delete(hands3)
	  d.Shuffle()
	  p.SetHand(hands,0,8)
	  p.SetHand(hands3,7,8)
	  skip:=[]int{0,7}
	  p.DealHand(d,skip)
	  tb.SetFlop(data)
	  fmt.Println(p.Index)
	  fmt.Println(len(d.Cards))
	   deal.Dealcard(p,data)
	   
	}