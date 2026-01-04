package card

type Flop interface {
	SetFlop(card [] int)
    SetTurn ( turn int)
	SetRiver(river int)
    Allcards() int
}
type Table struct{
	Flop [] int
	Turn int
	River int 
	All [] int
}
func (tb*Table)SetFlop(cards [] int) {
     tb.Flop =append(tb.Flop,cards...) 
}
func (tb*Table)SetTurn(turn int){
	tb.Turn=turn
}
func (tb*Table)SetRiver(river int){
	tb.River=river
}
func (tb*Table)Allcards()int{
	index:=0
	tb.All=append(tb.All,tb.Flop...)
    tb.All=append(tb.All,tb.Turn,tb.River)
	if len(tb.All)==5{
	 index =1
	}else{
		index=-1
	}
   return index
}