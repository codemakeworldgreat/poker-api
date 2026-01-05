package deals

import (
	"fmt"
	"pokers/card"
	"pokers/deals/dealcolors"
)
type CardsValue struct{
	Id int
	Card []int
	Color []int
	Number [] int
    
}

func Dealcard(p *card.Player,flopcard []int){
    value:=make([]CardsValue,len(p.Index),)
    
	for i:=0;i<len(p.Index);i++{
      value[i].Card=append(value[i].Card,p.Index[i].Hand...)
      value[i].Card=append(value[i].Card,flopcard...)
	  value[i].Id=i
	  deal:=value[i].Card

      for j,_:=range deal{
	  value [i].Number=append(value [i].Number,deal[j]>>2)
	  value[i].Color=append(value[i].Color,deal[j]&3)
	}
	fmt.Println(value[i])
     setcolor:=[]int{0,1,2,0,0,0,1}
    dealcolors.DealFlush(setcolor)
	
}}
