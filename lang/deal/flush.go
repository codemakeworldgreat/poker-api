package deal

import "fmt"

func DealFlush(color[] int,list[]int)[]int{
	 result:=make([]int,0)
   count:= DealsCount(color,4)
   fmt.Println(count)
   for i,v:=range count{
	if v>=5{
		index:=i
       result= findCards(index,list)
	    sort:=sortNumber (result)
		straights:=dealstraight(sort)
		if len(straights)>0{
			if straights[0]==14&&straights[4]==10{
              //最大同花顺
			}else{
             //同花顺
			}
		}
	   
	}else{
	   continue
	}
   }
   return result
}
func findCards(index int,list[] int)[]int{
	 flush:=make([]int,0,len(list))
     for _,v:=range list{
		if v&3==index{
         flush=append(flush,v>>2)
		}
	 }
	 return flush
}