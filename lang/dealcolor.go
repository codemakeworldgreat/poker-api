package main;

import "sort"
func dealcolor(point*Points,number []int,color[]int) {

count:=make([][]int,4)
 
for i,v:=range color{
	
count[v]=append(count[v],number[i])
}

 for i:=0;i<4;i++{
	if len(count[i])>=5{
		point.Flush=5
        flush:=count[i]
	sort.Sort(sort.Reverse(sort.IntSlice(flush)))
      point.Cards=flush[0:5]
	  return
	 }
	 
	}
 }
 