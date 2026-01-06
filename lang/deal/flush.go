package deal

import "fmt"

func DealFlush(color[] int,list[]int)[]int{
	 result:=make([]int,0)
   count:= DealsCount(color,4)
   fmt.Println(count)
   for i,v:=range count{
	if v>=5{
		index:=i
       result= FlushCards(index,list)
	   fmt.Println(result)
	}else{
	   continue
	}
   }
   return result
}
func FlushCards(index int,list[] int)[]int{
	 flush:=make([]int,0,len(list))
     for _,v:=range list{
		if v&3==index{
         flush=append(flush,v>>2)
		}
	 }
	 sortflush:=sortNumber(flush)
	 return sortflush
}