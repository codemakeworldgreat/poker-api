package deal

func  (c*Counts)isFour(number[]int){
	max:=0
	
    if c.Four!=0{
       if len(c.HighCards)>0{
		 max=c.HighCards[0]
	   }
	   if len(c.TwoPairs)>0{
		if c.TwoPairs[0]>max{
		 max=c.TwoPairs[0]
	   }
	}
	   if c.Three>0{
         if c.Three>max{
			max=c.Three
		 }
	   }
	  fourkind:=make([]int,0)
	  fourkind=append(fourkind,c.Four,max)
	}
}