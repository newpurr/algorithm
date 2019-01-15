package main

import (
	"fmt"
)

func main() {
	slic := []interface{}{1, 2, 3, 4, 5}
	fmt.Print(subSet(slic))
}

func subSet(arr []interface{}) [][]interface{} {
	var (
		// 给定集合的长度
		num = len(arr)
		// 所有子集个数为2的n次方,详情查看数学集合部分
		maxCount = 1 << uint(num)
		// 子集数组
		subArray [][]interface{}
		//formatStr string
		nullArr   []interface{}
	)

	subArray = append(subArray, nullArr)

	// 共有maxCount个子集
	for i := 1; i < maxCount; i++ {
		var tArr = nullArr
		// 获取当前子集的实际组合
		// 这里遍历数组是为了将当前外层循环的可能性转换成实际组合数组
		for j := 0; j <= num; j++ {
			if (i & (1 << uint(j))) != 0 {
				tArr = append(tArr, arr[j])
			}
		}
		subArray = append(subArray, tArr)
	}
	return subArray
}
