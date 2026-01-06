package deal



func  (c*Counts)isFour(number[]int)[]int{
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
	  fourkind=append(fourkind,8,c.Four,max)
	  return fourkind
	}
	return nil
}
func (c*Counts)isThree()[]int{
     data:=make([]int,0,3)
   if c.Three!=0{
	if len(c.TwoPairs)>0{
      data=append(data,7,c.Three,c.TwoPairs[0])
	  return data
	}else{
		data=append(data,4,c.Three,c.HighCards[0])
		return data
	}
   }
   return nil
}
func (c*Counts)isTwo()[]int{
	data:=make([]int,0)
	if len(c.TwoPairs)>1{
	 data=append(data,3,c.TwoPairs[0],c.TwoPairs[1],c.HighCards[0])
	 return data
	}else if len(c.TwoPairs)==1{
		data=append(data,2,c.TwoPairs[0],c.HighCards[0])
		return data
	}
	return nil
}
func (c*Counts)isHigh()[]int{
	result:=[]int{1,}
	result=append(result,c.HighCards...)
   return result
}