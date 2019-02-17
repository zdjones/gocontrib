package p

// prove should be able to eliminate bounds check in the for-loop here
// func f(b []byte, n int) {
// 	if n < 0 {
// 		panic("")
// 	}

// 	// _ = b[n]  // Bounds check
// 	_ = b[25] // Bounds check

// 	for i := n; i <= 25; i++ {
// 		b[i] = 123 // Bounds check
// 	}
// }

// // Should be enough with the first bounds check (?)
// func f2(b []byte, n int) {
// 	_ = b[n]  // Bounds check
// 	_ = b[25] // Bounds check

// 	for i := n; i <= 25; i++ {
// 		b[i] = 123 // Bounds check
// 	}
// }

// // when using uint, it works as expected and eliminates bounds check in the loop
// func foo3(b []byte, n uint) {
// 	_ = b[25] // Bounds check

// 	for i := n; i <= 25; i++ {
// 		b[i] = 123 // Bounds check elimintated
// 	}
// }

// func fdan(i int, data []byte) {
// 	if i >= 0 {
// 		for i < len(data) {
// 			_ = data[i] // bounds check!
// 			i++
// 		}
// 	}
// }

// func fslice(s []int) {
// 	if len(s) >= 2 {
// 		s = s[1:]
// 		_ = s[0]
// 	}
// }

// cmd/compile: prove pass disabled on int32/int64 #29964
// func f(x []int, j int) int {
// 	if x[j] != 0 {
// 		return 1
// 	}
// 	if j > 0 && x[j-1] != 0 {
// 		return 1
// 	}
// 	return 0
// }

func f32(x []int, j int32) int {
	if x[j] != 0 {
		return 1
	}
	if j > 0 && x[j-1] != 0 {
		return 1
	}
	return 0
}
