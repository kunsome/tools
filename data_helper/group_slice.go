package data_helper

// GroupSlice 把切片按照固定大小分组
// 参数：
// - arr 待切分数组
// - num 分组大小
// 返回值：
// - [][]T 切分好的数组
func GroupSlice[T any](arr []T, num int) [][]T {
	max := len(arr)
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num || num == 0 {
		return [][]T{arr}
	}
	//获取应该数组分割为多少份
	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]T, 0)
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
