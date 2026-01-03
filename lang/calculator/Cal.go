package  calculator
import"fmt"
 
 type Result struct{
    Frequency float32 
	Players []Player
	Person int
    Index []*Player
	
 }

 
func (result*Result)calculator(data*DataCard) {
	 
	 person:=data.Person
	results:=&Result{Frequency:float32(1),Person:person ,Index:make([]*Player,person)}
	
	 
	win:=float32(0)
	results.Players=make([]Player,results.Person)//初始化
	
	for i:=0;i<results.Person;i++{//指针索引初始化 多个对手
		results.Index[i]=&results.Players[i]
		results.Players[i].Hand=make([]int,0)
	}
	 
    deck:=&Deck{Cards:make([]int,52),Hands:make([]int,0)}
	deck.Hands=append(deck.Hands,data.Card1,data.Card2)
	
     deck.initial()//初始化牌组
	deck.know(data,results.Index,data.Card1,data.Card2);//删掉已知牌 和初始化写一块了
    
    copycard:=make([]int,len(deck.Cards))
	copy(copycard,deck.Cards)
	 
	
    pre:=float32(100000)
	for i:=0;i<int(pre);i++{
    deck.shuffle()//洗牌
	
	deck.deal(results.Index,results,data)//发牌
	
	win +=dealpoint(deck,results.Index)//算牌逻辑
	
	deck.Hands=deck.Hands[:2]//还原手牌
	
	for j:=0;j<results.Person;j++{
     
	  results.Index[j].Hand=results.Index[j].Hand[:0]  
	}
	
    deck.Cards=deck.Cards[:0]
	deck.Cards=append(deck.Cards,copycard...)
 } 
    
   
  win32:=float32(win)
   cal:=(win32/pre)*100
   fmt.Println(win32,cal)

}

