package main

import "fmt"

//二分查找
func Binaryfind(arr *[6]int,left int,right,findVal int)  {

	if left>right{
		fmt.Println("找不到")
	}
	mid :=(left+right)/2
	if arr[mid]>findVal{
		Binaryfind(arr,left,mid-1,findVal)
	}else if arr[mid]<findVal {
		Binaryfind(arr,mid+1,right,findVal)
	}else {
		fmt.Println("找到了" )
	}
	fmt.Println(*arr)
}

func main() {
	arr :=[6]int{1,2,3,57,87,222}
	fmt.Println(arr)
	Binaryfind(&arr)

}
