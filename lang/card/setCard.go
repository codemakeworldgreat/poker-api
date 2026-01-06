package card

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

func InitPlayer(person int)*Player{
	player:=&Player{Index:(make([]Detail,person)),
	Person:person}
	return player
}
func (p*Player)SetHand(hand[]int,id int){
   
 for i:=0;i<p.Person;i++{
   p.Index[i].Id=i
 }
 p.Index[id].Hand=hand 
}
func (p*Player)DealHand(d*DeckCard, skipId int){
   
	for i:=0;i<p.Person;i++{
		skip:=false
		 
			if i==skipId{
              skip=true
			  break
			}
        if skip{
			continue
		}
      next:= d.Send(2)
	  p.Index[i].Hand=append(p.Index[i].Hand,next...)
	}
}