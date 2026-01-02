package main
import ("sort"

)
func sortdeal(number[]int,level*Points){
	level.Cards=make([]int,0)
	if level.Level==4{
      if level.Maxcard==5{
       level.Cards=append(level.Cards,5,4,3,2,14)
	  }else{
		for j:=0;j<5;j++{
			level.Cards=append(level.Cards,level.Maxcard-j)
		}
	  }
	} else{
     if level.Level==7{
		level.Cards=sortnumber(number,level.Maxcard)
		
}else if level.Level==6{
      level.Cards=sortnumber(number,level.Maxcard,level.First)
	 }else if level.Level==3{
		level.Cards=sortnumber(number,level.Maxcard)
	 }else if level.Level==2{
		level.Cards=sortnumber(number,level.First,level.Second)
	 }else if level.Level==1{
		level.Cards=sortnumber(number,level.Maxcard)
	 }else{
		level.Cards=sortnumber(number,level.Maxcard)
	 }
}
}
func sortnumber(number[]int,index...int)[]int{
	
	sortup:=make([]int,len(number))
	list:=make([]int,0)
	copy(sortup,number)

	sort.Slice(sortup,func(i,j int)bool{
		return sortup[i]>sortup[j]
	})
	for _,v:=range sortup{
		for _,j:=range index{
           if v==j{
            list=append(list,v)
		   }
	}
	}
	for _,v:=range sortup{
      isindex:=false 
	   for _,j:=range index{
		 if v==j{
			isindex=true
		 }
	   }
	   if isindex==false{
        list=append(list,v)
	   }
	}
	if len(list)>5{
		list=list[:5]
	}
	return list
}