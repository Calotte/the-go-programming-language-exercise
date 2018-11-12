package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	cnt := 0
	for x != 0 {
		cnt += int(pc[byte(x)])
		x >>= 8
	}
	return cnt
}
