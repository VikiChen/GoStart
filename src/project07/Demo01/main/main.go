package main

import "fmt"

//冒泡排序
func BubbleSort(arr *[5]int)  {
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j< len(arr)-1-i;j++{
			if arr[j]>arr[j+1]{
				temp := arr[j+1]
				arr[j+1]=arr[j]
				arr[j]=temp
			}

		}
	}

	fmt.Println(*arr)
}

func main() {
	arr :=[5]int{11,123,4,577,87}
	fmt.Println(arr)
	BubbleSort(&arr)

}
