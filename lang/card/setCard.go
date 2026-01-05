package card
import "fmt"
type Detail struct{
	Id int 
	Hand []int
	Color [] int
    Number [] int
    score int
}
type Player struct{
	Index [] Detail
    Person int
}

func (p*Player)SetHand(hand[]int,id int,person int){
   p.Person=person
 for i:=0;i<person;i++{
   p.Index[i].Id=i
 }
 p.Index[id].Hand=hand 
}
func (p*Player)DealHand(d*DeckCard, skipId[]int){
   fmt.Println( len(d.Cards))
	for i:=0;i<p.Person;i++{
		skip:=false
		for _,v:=range skipId{
			if i==v{
              skip=true
			  break
			}
		
		}
        if skip{
			continue
		}
      next:= d.Send(2)
	  p.Index[i].Hand=append(p.Index[i].Hand,next...)
	}
}