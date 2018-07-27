package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr)<2 {
		return arr
	}else{
		prov:=arr[0]
		arr=arr[1:]
		less:=filter(arr, func(i int, v int) bool{
			return v<=prov
		})
		greater:=filter(arr, func(i int,v int) bool {
			return v>prov
		})
		less=quickSort(less)
		greater=quickSort(greater)
		less=append(less, prov)
		less= append(less, greater...)
		return less
	}
}


func filter(arr []int, fn func(index int, value int) bool ) []int {
	result:=make([]int, 0)
	for i,v := range arr {
		if flag:=fn(i,v); flag{
			result = append(result, v)
		}
	}
	return result
}

func main() {
	arr:=[]int{-1,2,99,101,5,1,2,45,3,4,5,6}
	arr= quickSort(arr)
	fmt.Println(arr)
}