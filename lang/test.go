package main
import "fmt"
func test(){
	heroNumber := []int{11, 8, 10, 10, 14, 9, 13} 
	    heroColor  := []int{ 2,  3, 0, 0,  0,  1,  1} 

		    // 2. 构造 Villain (同花顺)
			    // Villain 手牌: 2(方片), 3(方片), 4(方片), 5(方片), A(方片), Q(梅花), K(梅花)
				    // 这就是 A-2-3-4-5 同花顺 (Level 8)
					    villainNumber := []int{2, 2, 5, 9, 4, 6, 14}
						    villainColor  := []int{0, 2, 3, 0,  2,  1,  1}
							points1:=dealnumber(heroNumber,heroColor)
							points2:=dealnumber(villainNumber,villainColor)
							fmt.Println(points1)
							fmt.Println(points2)

}