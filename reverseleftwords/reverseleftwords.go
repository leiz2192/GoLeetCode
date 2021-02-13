package reverseleftwords

// ReverseLeftWords https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func ReverseLeftWords(s string, n int) string {
	b := []byte(s)
	rst := b[n:]
	rst = append(rst, b[:n]...)
	return string(rst)
}
