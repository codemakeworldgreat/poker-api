package deal

import "fmt"

func DealFlush(color[] int,list[]int)[]int{
	 find:=make([]int,0)
	result:=make([]int,0)
   count:= DealsCount(color,4)
   fmt.Println(count)
   for i,v:=range count{
	if v>=5{
		index:=i
       find= findCards(index,list)
	    sort:=sortNumber (find)
		straights:=dealstraight(sort)
		if len(straights)>0{
			if straights[0]==14&&straights[4]==10{
              result=append(result,10)
			  result=append(result,straights...)
			  result=append(result,i)
			  return result
			}else{
			   result=append(result,9)
			   result=append(result,straights...)
			   result=append(result,i)
             //同花顺
			 return result
			}
		}
          //同花
              result=append(result,6)
	          result=append(result,sort...)
			  result=append(result,i)
			  return result
	}else{
	   continue
	}
   }
   return nil
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