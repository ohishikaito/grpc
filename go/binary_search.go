package main

// func main() {
// 	slice := []int{}
// 	for i := 0; i < 100; i++ {
// 		slice = append(slice, i)
// 	}

// 	// var target int
// 	// fmt.Scanf("%d", &target)
// 	target := 48

// 	count := 0
// 	index, count := binarySearch(slice, 0, len(slice)-1, target, count)
// 	if index < 0 {
// 		fmt.Println("target not found")
// 	} else {
// 		fmt.Printf("index: %d, count: %d", index, count)
// 	}
// }

// // [1, 2..98, 99], 0, (100-1)=99, 19, 0 =>
// // -1 を返すときは、異常値っぽい
// func binarySearch(slice []int, low int, high int, target int, count int) (int, int) {
// 	fmt.Println("low", low)
// 	fmt.Println("high", high)
// 	// 0以下 or targetが配列の先頭より小さい or targetが配列の最大より大きい とreturn
// 	if len(slice) < 1 || target < slice[0] || slice[high] < target {
// 		return -1, count
// 	}

// 	mid := (low + high) / 2 // 中央値を定義 (0+99)/2 => 49
// 	count++                 // 比較回数をプラス

// 	// 中央値とtargetが一緒ならreturn
// 	if slice[mid] == target {
// 		return mid, count
// 	}

// 	// これ使うのか？これ以上探索するデータないらしいけど
// 	if low >= high {
// 		return -1, count
// 	}

// 	// 中央値がtargetより大きければ、high(上限)を中央値-1(48)にして再帰
// 	if slice[mid] > target {
// 		return binarySearch(slice, low, mid-1, target, count)
// 	}

// 	// 中央値がtargetより小さければ、high(上限)を最低値をmid+1(50)と上限にして再帰
// 	if slice[mid] < target {
// 		return binarySearch(slice, mid+1, high, target, count)
// 	}

// 	return -1, count
// }
