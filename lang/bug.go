package main 
import "fmt"
func bug (deck  []int,p[]*Player){
	array:=make([]int,0)
	array=append(array,deck...)
	array1:=make([]int,0)
	for i:=0;i<8;i++{
      array1=append(array1,p[i].Hand[:2]...)
	}
	fmt.Println(array1)
	fmt.Println(array)
}