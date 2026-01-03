package card

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
func setMap(hands[]string,publicCard []string)([]int,[]int){
	var flop[] int
	var hand[]int
	maplook:=getMap()
	 for _,v:=range hands{
		val,ok:=maplook[v]
	   if ok{
          hand=append(hand,val)
	   }   
	 }
	if len(publicCard)!=0{
		for _,v:=range publicCard{
			val,ok:=maplook[v]
			if ok{
			 flop=append(flop,val)
			 }
		}
	}    
     return hand,flop
}

func countcard() map[int]int{
	cards:=getMap()
	countMap:=make(map[int] int)
	 for _,v:=range cards{
       countMap[v]=0
	 }
   return countMap
}