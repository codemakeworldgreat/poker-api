package card


type PublicCards struct{
	Flop [] int
	
}
func setpublic (card ...int) []int{
     p:=PublicCards{}
	 p.Flop=append(p.Flop,card...)
   return p.Flop 
}