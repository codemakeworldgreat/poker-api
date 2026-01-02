package main

import (
	
	"strconv"
	)
func getMap() map[string]int{
	color:=make(map[string]int)
		color["S"]=0
         color["H"]=1
		 color["D"]=2
		 color["C"]=3
	
    number:=make(map[string]int)
    for i:=2;i<=10;i++{
       number[strconv.Itoa(i)]=i
	}
    number["J"]=11
	number["Q"]=12
	number["K"]=13
    number["A"]=14
	
    bitCard:=make(map[string]int)
	for n,v:=range number{
    for z,x:=range color{
		key:=n+z
	 bitCard[key]= (v<<2)|x
	}
} 
  
   return bitCard
}
func setMap(card1,card2 string,publicCard []string)([]int,int, int){
	

	var bitwise[] int
	maplook:=getMap()
	if len(publicCard)!=0{
		for _,v:=range publicCard{
			val,ok:=maplook[v]
			if ok{
			 bitwise=append(bitwise,val)
			    
			}else {
				
			}
		}
	}
    bitwise1:=maplook[card1]
    bitwise2:=maplook[card2]
	return bitwise,bitwise1,bitwise2
}