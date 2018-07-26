package main

import "fmt"

func sum(arr []int) int {
	if len(arr) ==0 {
		return 0
	}else if len(arr)==1{
		return arr[0]
	}else{
		x:=arr[0]
		return x+sum(arr[1:])
	}
}

func count(arr []int) int {
	if len(arr)==0 {
		return 0
	}else{
		return 1+count(arr[1:])
	}
}

func max(arr []int) int {
	if len(arr) ==0{
		return 0
	}else if len(arr)==1{
		return arr[0]
	}else{
		biggest:=arr[0]
		if biggest>max(arr[1:]) {
			return biggest
		}else{
			return max(arr[1:])
		}
	}
}


func main(){
	arr:=[]int{1,2,9,4,5}
	fmt.Printf("sum=%d\n",sum(arr))
	fmt.Printf("count=%d\n",count(arr))
	fmt.Printf("max=%d\n",max(arr))
}