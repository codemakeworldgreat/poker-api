package card
import "strconv"
type CardsMap interface{
   SetMap(hand[] string)[]int
}


func (d*DeckCard) getMap() {
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
	
   
	for n,v:=range number{
    for z,x:=range color{
		key:=n+z
	 d.MapCards[key]= (v<<2)|x
	}
} 
  
}
func (d*DeckCard)SetMap(hands[]string)([]int){
	 
	var hand[] int
	maplook:=d.MapCards
	 for _,v:=range hands{
		val,ok:=maplook[v]
	   if ok{
          hand=append(hand,val)
	   }   
	 }
   
     return hand
}

