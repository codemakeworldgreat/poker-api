package card


type Playerdetail struct{
	Id int 
	Hand []int
	Color [] int
    Number [] int
}
type Player struct{
	Players [] Playerdetail

}

func Setcard(hand[]int,id int,Person int)*Player{
   player:=&Player{ Players:make([]Playerdetail,Person),   
}
   player.Players[id].Id=id
   player.Players[id].Hand=hand
 
    for i:=0;i<len(hand);i++{
      player.Players[id].Color[i]=hand[i]&3
      player.Players[id].Number[i]=hand[i]>>2
    }
   return player
}