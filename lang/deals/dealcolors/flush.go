package dealcolors
import "pokers/deals"
import "fmt"
func DealFlush(color[] int){
	
   count:= deals.DealsCount(color,4)
   for i,v:=range count{
	if v==5{
		fmt.Println("同花",i)
	}
   }
}