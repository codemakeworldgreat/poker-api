package main;
import "fmt"

import "poker/card"
type DataCard struct{
	Card1 int
	Card2 int
	PublicCard []int
	Person int 
    
}

 
func main(){
  
 // requestcard();//启动接收数据从laravel
       
     
    
	  tb:=card.Table{}
      d:=card.GetInitial()
	 public:=[]string{"AC","6H","5C","10H","8S"}
      data:=d.SetMap(public)
	  hand:=[]string{"8H","8S"}
      hands:=d.SetMap(hand)
	  fmt.Println(data,hands)
	  d.Delete(hands)
	  d.Delete(data)
	  d.Shuffle()
	  
	  tb.SetFlop(data[:3])
        turns:=d.Send(1)
		tb.SetTurn(turns[0])
	   fmt.Println(tb.Turn)
	}