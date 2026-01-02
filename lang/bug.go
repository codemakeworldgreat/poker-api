package main 
import "fmt"
 var countMap map[int]int
func bug (deck  []int,p[]*Player){
      
	 countMap= countcard()
	array:=make([]int,0)
	array=append(array,deck...)
	array1:=make([]int,0)
	for i:=0;i<8;i++{
      array1=append(array1,p[i].Hand...)
	}
   fmt.Println(array1)
   
}