package deal

import "fmt"

func DealFlush(color[] int,list[]int)int{
	 result:=-1
   count:= DealsCount(color,4)
   fmt.Println(count)
   for i,v:=range count{
	if v>=5{
		result=i
       return result
	}else{
	   continue
	}
   }
   return result
}
func FlushCards(index int,list[] int){
	 flush:=make([]int,0,len(list))
     for _,v:=range list{
		if v&3==index{
         flush=append(flush,v>>2)
		}
	 }
}