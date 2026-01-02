package main;

func dealFlush(points *Points) {
 if points.Flush==5{
  if points.Level==4{
    if points.Maxcard==14{//最大的同花顺
      points.Level=9
    }else{
      points.Level=8
      points.Maxcard=points.Cards[0]
    }
  }else{
    if 5>points.Level{
      points.Level=5
      points.Maxcard=points.Cards[0]
    }
  } 
 }
}