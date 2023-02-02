package rbtree

import "testing"

func TestRBTree_Insert(t *testing.T) {
	tests := []struct {
		idsToInsert    []int
		idsToFind      []int
		correctAnswers []bool
	}{
		// test 01
		{
			idsToInsert:    []int{0},
			idsToFind:      []int{1, 0},
			correctAnswers: []bool{false, true},
		},
		//test 02
		{
			idsToInsert:    []int{20, 10, 30},
			idsToFind:      []int{5, 10, 15, 20, 25, 30, 35},
			correctAnswers: []bool{false, true, false, true, false, true, false},
		},
		//test 03
		{
			idsToInsert:    []int{10, 20, 30},
			idsToFind:      []int{5, 10, 15, 20, 25, 30, 35},
			correctAnswers: []bool{false, true, false, true, false, true, false},
		},
		//test 04
		{
			idsToInsert:    []int{30, 20, 10},
			idsToFind:      []int{5, 10, 15, 20, 25, 30, 35},
			correctAnswers: []bool{false, true, false, true, false, true, false},
		},
		//test05
		{
			idsToInsert: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 20, 30, 40, 50,
				60, 70, 80, 90, 100},
			idsToFind: []int{1, 10, 2, 20, 3, 30, 5, 50, 8, 80, 13, 130},
			correctAnswers: []bool{true, true, true, true, true, true, true,
				true, true, true, false, false},
		},
	}
	for i := range tests {
		answers := make([]bool, 0, len(tests[i].correctAnswers))
		rbt := NewRBTree()
		for _, x := range tests[i].idsToInsert {
			rbt.InsertNode(x)
		}
		for _, x := range tests[i].idsToFind {
			_, err := rbt.GetTime(x)
			answers = append(answers, err == nil)
		}
		for j, x := range answers {
			if x != tests[i].correctAnswers[j] {
				t.Errorf("Wrong test %d of %d", i+1, len(tests))
			}
		}
	}

}
