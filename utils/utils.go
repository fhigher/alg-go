package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 计算f运行时间
func RunningTime(funcName string, f func(arr []int, n int), arr []int, n int) {
	//fmt.Println("before sorting: ", arr)
	startTime := time.Now().UnixNano()
	f(arr, n)
	endTime := time.Now().UnixNano()
	fmt.Printf("%s running %f seconds\n", funcName, float64(endTime-startTime)/1000/1000/1000)
	//fmt.Println("after sorting: ", arr)
}

// 生成随机数数组
func GenerateRandNums(n int) (arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(n))
	}
	return
}
