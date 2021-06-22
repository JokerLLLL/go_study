package algo

import (
	"testing"
)

func TestSearchLengthSonString(t *testing.T) {
	tests := []struct {
		a string
		b int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"abc", 3},
		{"h;l;l朝陈", 4},
		{"", 0},
		{"pwkew1234", 7},
	}
	for _ , test:= range tests {
		if actual:= SearchLengthSonString(test.a); test.b != actual {
			t.Errorf("test SearchLengthSonString ERROR: use a:%s should get b:%d, actual:%d", test.a, test.b, actual)
		}
	}

}

func BenchmarkSearchLengthSonString(b *testing.B) {
	a := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i:=0;i<13;i++ {
		a += a
	}
	b.Logf("a len:%d", len(a))
	ans := 8
	b.ResetTimer() //准备数据时间不算
	for i := 0; i < b.N; i++ {
		actual := SearchLengthSonString(a)
		if ans != actual {
			b.Errorf("test SearchLengthSonString ERROR: use a:%s should get b:%d, actual:%d", a, ans, actual)
		}
	}
}