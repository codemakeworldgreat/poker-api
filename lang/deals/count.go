package deals

func DealsCount(data[]int,size int)[]int{
   count:=make([]int,size)
   for _,v:=range data{
	  count[v]++
   }
   return count 
}

