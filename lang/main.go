package main

import (
	"poker/card" 
)
type DataCard struct{
	Card1 int
	Card2 int
	PublicCard []int
	Person int 
    
}

 
func main(){
  
 // requestcard();//启动接收数据从laravel
       //设置手牌公牌删除 
    mycard:=[]string{"AC","8H"}
	public:=[]string{"AD","10S","QC"}
     card.Requestcard(mycard,public,9)
	}