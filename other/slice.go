package main

import (
	"fmt"
	"strconv"
)

func splitArray(arr []NewNum, num int) [][]NewNum {
	//声明分割好的二维数组
	var segments = make([][]NewNum, 0)

	max := int(len(arr))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		segments = append(segments, arr)
		return segments
	}
	//获取应该数组分割为多少份
	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割数组的截止下标
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i*num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i*num
	}
	return segments
}


type NewNum struct {
	Num int
}

func main() {
	arr1:=[]NewNum{}
	//数组0~24
	for i := 0; i < 25; i++ {
		arr1=append(arr1, NewNum{i})
	}
	//拆分4份
	fmt.Println("println:",splitArray(arr1,4))
	//拆分7份
	fmt.Println("println:",splitArray(arr1,7))
	//拆分12份
	fmt.Println("println:",splitArray(arr1,2))

	var s [][]OaidRequestItem
	var sr [][]OaidRequestItem
	list := make([]OaidRequestItem, 1)
	for i :=range list{
		list[i] = OaidRequestItem{strconv.Itoa(i)}
	}
	s = append(s, list)
	fmt.Println(s)

	for _,row :=range s {
		s2 := splitArray2(row, 3)
		for _,item := range s2 {
			sr = append(sr, item)
		}
	}

	fmt.Println(sr)

}

type OaidRequestItem struct {
	Info string
}

func splitArray2(arr []OaidRequestItem, num int) [][]OaidRequestItem {
	var segments = make([][]OaidRequestItem, 0)

	max := int(len(arr))
	if max <= num {
		segments = append(segments, arr)
		return segments
	}
	//获取应该数组分割为多少份
	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割数组的截止下标
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}