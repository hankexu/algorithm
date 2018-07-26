package main

import "fmt"

func binarySearch(list []int , key int) int {
	low,high:=0, len(list)-1
	for low<=high {
		mid:=(high+low)/2
		guess:=list[mid]
		switch {
		case guess>key: high = mid -1
		case guess<key: low = mid+1
		case guess==key: return mid
		}
	}
	return -1
}

func main() {
  a:=[]int{1,2,3,4,5,6,7,8,9,10}
  index:=binarySearch(a, 5)
  fmt.Println(index)
}