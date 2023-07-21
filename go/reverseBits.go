package main

// 方法二 分治
const (
	m1 = 0x55555555 // 01010101010101010101010101010101
	m2 = 0x33333333 // 00110011001100110011001100110011
	m4 = 0x0f0f0f0f // 00001111000011110000111100001111
	m8 = 0x00ff00ff // 00000000111111110000000011111111
)

func reverseBits(n uint32) uint32 {
	n = n>>1&m1 | n&m1<<1
	n = n>>2&m2 | n&m2<<2
	n = n>>4&m4 | n&m4<<4
	n = n>>8&m8 | n&m8<<8
	return n>>16 | n<<16
}

// 方法一 逐位颠倒
//func reverseBits(n uint32) uint32 {
//	var ans uint32 = 0
//	for i := 0; i < 32 && n > 0; i++ {
//		ans |= (n & 1) << (31 - i)
//		n >>= 1
//	}
//	return ans
//}

// func main() {
// 	var n uint32 = 31243
// 	ret := reverseBits(n)
// 	fmt.Println(ret)
// }
