package deals
func  Bitwise(cards[]int)([]int,[]int){
	number:=make([]int,len(cards))
    color:=make([]int,len(cards))
	for i,v:=range cards{
        number[i]=v>>2
		color[i]=v&3
	}
	return number,color
}