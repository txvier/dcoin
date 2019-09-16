package main

import (
	"fmt"
	"time"
)

func findNumsByIndexs(nums []string, indexs [][]int) [][]string {
	if len(indexs) == 0 {
		return [][]string{}
	}

	result := make([][]string, len(indexs))
	for i, v := range indexs {
		line := make([]string, 0)
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}
		}
		result[i] = line
	}
	return result
}

func addTo(arr [][]int, ele []int) [][]int {
	newEle := make([]int, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)
	return arr
}

func moveOneToLeft(leftNums []int) {
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}
	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

func jieCheng(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func mathZuhe(n, m int) int {
	return jieCheng(n) / (jieCheng(n-m) * jieCheng(m))
}

func zuheResult(n, m int) [][]int {
	if m < 1 || m > n {
		fmt.Println("enen")
		return [][]int{}
	}

	result := make([][]int, 0, mathZuhe(n, m))

	indexs := make([]int, n)

	for i := 0; i < n; i++ {
		if i < m {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}

	result = addTo(result, indexs)
	for {
		find := false

		for i := 0; i < n-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true
				indexs[i], indexs[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexs[:i])
				}
				result = addTo(result, indexs)
				break
			}
		}
		if !find {
			break
		}
	}
	return result
}

func all(a []string, c int) {
	first := a[0]
	l := len(a)
	indexs := zuheResult(l, c)
	result := findNumsByIndexs(a, indexs)
	for _, v := range result {
		p := append(v, first)
		fmt.Println(p)
	}
}

func main() {
	arr := []string{"1", "2", "3", "4"}
	timeStart := time.Now()
	for i := 3; i <= len(arr); i++ {
		all(arr, i)
	}
	timeEnd := time.Now()
	//fmt.Println("result:", result)
	fmt.Println("time consume:", timeEnd.Sub(timeStart))

	//rightCount := mathZuhe(n,m)
	//if rightCount==len(result){
	//	fmt.Println("right")
	//}else{
	//	fmt.Println("error,the right result is:",rightCount)
	//}
}
