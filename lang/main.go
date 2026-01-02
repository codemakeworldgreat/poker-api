package main;
import "fmt"
type DataCard struct{
	Card1 int
	Card2 int
	PublicCard []int
	Person int 

}
func main(){
  
 // requestcard();//启动接收数据从laravel
       
     result:=&Result{}
     //var public [] string
	 public:=[]string{}

	 set,card1,card2:=setMap("AH","AD",public)
	 data:=&DataCard{Card1:card1,Card2:card2,PublicCard:set,Person:8}
	 
     
	 result.calculator(data)
	 fmt.Println(count)
	 
	 fmt.Println("lose",lose)
	 fmt.Println(countMap)
	}