package deal
import "fmt"
type Counts struct{
	Four int
    Three int
	TwoPairs []int
	HighCards[] int
}
 
func CountNumber(number[]int)[]int{
	count:=DealsCount(number,15)
     fmt.Println(count)
    counts:=&Counts{TwoPairs:make([]int,0)}
	
  for i:=len(count)-1;i>0;i--{
	v:=count[i]
	switch v{
	case 4:
	 counts.Four=i
	case 3:
	  counts.Three=i
	case 2:
	counts.TwoPairs=append(counts.TwoPairs,i)
	case 1:
	counts.HighCards=append(counts.HighCards,i)
    }
}
  fmt.Println(counts.HighCards)
 return count
}