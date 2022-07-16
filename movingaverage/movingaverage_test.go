package movingaverage_test

import (
	"testing"
)

// https://leetcode.cn/problems/qIsx9U/

type MovingAverage struct {
	items []int
	size  int
	sum   int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		items: make([]int, 0, size),
		size:  size,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	m.items = append(m.items, val)
	m.sum += val

	if m.size >= len(m.items) {
		return float64(m.sum) / float64(len(m.items))
	}
	m.sum -= m.items[0]
	m.items = m.items[1:]
	return float64(m.sum) / float64(m.size)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func TestMovingAverage(t *testing.T) {
	IsEqual := func(lhs, rhs float64) bool {
		if lhs > rhs {
			return lhs - rhs < 0.000001
		}
		return rhs - lhs < 0.000001
	}

	m := Constructor(3)
	got := m.Next(1)
	if !IsEqual(1.0, got) {
		t.Fatalf("expect: 1.0, got: %f", got)
	}

	got = m.Next(10)
	if !IsEqual(5.5, got) {
		t.Fatalf("expect: 5.5, got: %f", got)
	}

	got = m.Next(3)
	if !IsEqual(4.666667, got) {
		t.Fatalf("expect: 4.666667, got: %f", got)
	}

	got = m.Next(5)
	if !IsEqual(6.0, got) {
		t.Fatalf("expect: 6.0, got: %f", got)
	}
}
