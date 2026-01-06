package card
import "fmt"
func Requestcard(hand[]string,public[]string,person int){
	// 过map  从牌堆清空已知牌
	init:=GetInitial()//初始化
	myhand:=init.SetMap(hand)
     publics :=init.SetMap(public)
	 pool:=make([]int,0)
	 pool=append(pool,myhand...)
	 pool=append(pool,publics...)
	init.Delete(pool)
	//设置牌组 设置需要跳过的牌组 把每个对手两张牌补上
    player:=InitPlayer(person)
	player.SetHand(myhand,0)
	skip:=0
	player.DealHand(init,skip)
    fmt.Println(player)
}