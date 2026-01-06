package deal

import (
	"fmt"
	"poker/card"
	
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
     }
     test:=[]int{3,1,2,3,3,3,3,3}
	result:= DealFlush(test)
    fmt.Println(result)
	testnumber:=[]int{8,11,11,14,9,9}
	numbercount:=CountNumber(testnumber)
	fmt.Println(numbercount)
	sort:=sortNumber(value[1].Number)
	fmt.Println(sort)
}
