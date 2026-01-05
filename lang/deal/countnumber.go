package deal
import "fmt"
type Counts struct{
	Four int
    Three int
	TwoPairs []int
	
}
 
func CountNumber(number[]int)[]int{
	count:=make([]int,15)
    
 for _,v:=range number{
   count[v]++
 }
    counts:=&Counts{TwoPairs:make([]int,0)}
	
  for i,v:=range count{
	switch v{
	case 4:
	 counts.Four=i
	case 3:
	  counts.Three=i
	case 2:
	counts.TwoPairs=append(counts.TwoPairs,i)
    }
}
  fmt.Println(counts)
 return count
}