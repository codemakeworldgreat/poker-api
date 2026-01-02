package main;

type Straight struct{
	Point int 
	Max int
	Number []int
}
type Points struct{
	Flush int//同花的计数
	Level int
	Maxcard int
	FlushCards int 
	Cards []int 
	First int
	Second  int
}
type Fullhouse struct{
	Three  int
	Two int
}

func dealnumber(number[]int,color[]int)Points{
	
	level:=Points{}
	full:=Fullhouse{}
	
    stra:=Straight{Number:number};
	stra.straight()
  three:=false;
   threemax:=0
 two:=false;
 firstpair:=0
 secondpair:=0
 pairs:=0;
 onemax:=0
 count:=[15]int{};
 for _,v:=range number{
	count[v]++;
 }
 for i:=14;i>=2;i--{
    if (count[i]==4){	
     level.Level=7
	 level.Maxcard=i;
    
 } else if(count[i]==3){
	 if three==false{
		three=true
        threemax=i
	 } else{
      two=true
	if firstpair==0{
      firstpair=i
}
}
	  
 }else if(count[i]==2){
    two=true; 
	pairs++
	if firstpair==0{
		firstpair=i
	}else if secondpair==0 {
        secondpair=i
	}
}  else if count[i]==1{
	 
	if i>onemax{
		onemax=i
	}
	   } 	
}
  if level.Level==7{
  //确保四条没问题
  }else if(three==true&&two==true){
	 level.Level=6
	 full.Three=threemax
	 full.Two=firstpair
	 level.Maxcard=threemax
	 level.First=firstpair
  } else if stra.Point==4{//判断顺子
	  level.Level=4
	  level.Maxcard=stra.Max
  }else if three==true&&two==false{
      level.Level=3
	  level.Maxcard=threemax
      
  } else if three==false&&two==true{
	 if pairs>=2{
       level.First= firstpair
	   level.Second=secondpair
	   level.Level=2
	   level.Maxcard=firstpair
}else if pairs==1{
		level.Level=1
		level.Maxcard=firstpair
}}else{
	level.Level=0
	level.Maxcard=onemax
}
     
     sortdeal(number,&level)
    dealcolor(&level,number,color)
   dealFlush(&level)//最终分数
  
   
   return level  ;
}
func(st*Straight) straight(){
	 var  mask uint32=0
  for _,v:=range st.Number{
     value:=v-1
	 mask=mask|(1<<uint32(value))
  }
  if(mask>>13)&1==1{
	mask|=(1<<0)
  }
   s:=mask
   s&=(s<<1)
   s&=(s<<1)
   s&=(s<<1)
   s&=(s<<1)
   if s!=0{
	 st.Point=4
     var minkind int
	 for i:=14;i>0;i--{
		if (s>>uint(i))&1==1{
			minkind =i
			st.Max=minkind+1
			return
		}
       
	 }
  
   }
  
}