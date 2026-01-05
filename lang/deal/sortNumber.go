package deal

import "sort"

func sortNumber(number[]int)[]int{
    sort.Slice(number,func(i,j int)bool{
		return number[i]>number[j]
	})
	return number
}
