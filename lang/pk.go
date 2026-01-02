package main;

 var count=0
 
 var lose=0
var tie=0
func pk (mypoint*Points,player []*Player ) (float32){ 
  var win =0
   for i:=0;i<len(player);i++{
     index:=dealopp(mypoint,&player[i].Points)
     
     if index==-1{
      return 0
     }else if index==1{
       win++
     }else {
        tie++
}
}
 
  if win ==len(player){
    count++
    return 1
  }
return 0
}



func dealopp( p1, p2*Points)int{
 if p1.Level>p2.Level{
  return 1
 } else if p1.Level<p2.Level{
  return -1
 }else{
    for i:=0;i<5;i++{
      if p1.Cards[i]>p2.Cards[i]{
        
        return 1
      }else if p1.Cards[i]<p2.Cards[i]{
       return -1
      }
    }
  }
 return 0
}