package deal

import "fmt"
func DealFlush(color[] int)int{
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