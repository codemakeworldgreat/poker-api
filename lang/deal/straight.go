package deal
func dealstraight(number[] int)[]int{
	for i:=0;i<=len(number)-5 ;i++{
		if number[i]-4==number[i+4]{
		   return number[i:i+5]
		}
	}
	copynumber:=make([]int,0,5)
    if number[0]==14{
	for i:=1;i<=len(number)-4;i++{
      if number[i]==5&&number[i+3]==2{
        copynumber=append(copynumber,number[i:i+4]...)
		copynumber=append(copynumber,number[0])
		return copynumber
	  }
	}
	}
    return nil
}