package main

import "fmt"

func findSmallest(arr[]int) int{
	smallest,smallestIndex:=arr[0],0
	for i,val:=range arr{
		if val<=smallest {
			smallestIndex = i
			smallest = val
		}
	}
	return smallestIndex
}

func selectSort(arr []int) []int {
	list:=make([]int, 0)
	for len(arr)>0{
		smallestIndex:=findSmallest(arr)
		list = append(list, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return list
}


func main() {
  arr:=[]int{5,6,7,8,3,2,0,-1,5,3}

  list:= selectSort(arr)
  fmt.Println(list)
}