package rbtree

import "testing"

func TestRBTree_DeleteNode(t *testing.T) {
	tests := []struct {
		idsToInsert    []int
		idsToDelete    []int
		idsToFind      []int
		correctAnswers []bool
	}{
		// test 01
		{
			idsToInsert:    []int{0},
			idsToDelete:    []int{0},
			idsToFind:      []int{0},
			correctAnswers: []bool{false},
		},
		// test 02
		{
			idsToInsert:    []int{5, 10, 15, 20, 25, 30, 35},
			idsToDelete:    []int{20},
			idsToFind:      []int{20, 10, 30},
			correctAnswers: []bool{false, true, true},
		},
		// test 03
		{
			idsToInsert:    []int{5, 10, 15, 20, 25, 30, 35},
			idsToDelete:    []int{10, 30},
			idsToFind:      []int{5, 10, 15, 20, 25, 30, 35},
			correctAnswers: []bool{true, false, true, true, true, false, true},
		},
		// test 04
		{
			idsToInsert:    []int{5, 10, 15, 20, 25, 30, 35},
			idsToDelete:    []int{5, 15, 25, 35},
			idsToFind:      []int{5, 10, 15, 20, 25, 30, 35},
			correctAnswers: []bool{false, true, false, true, false, true, false},
		},
		// test 05
		{
			idsToInsert: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 20, 30, 40, 50,
				60, 70, 80, 90, 100},
			idsToDelete: []int{5, 10, 40, 7, 80, 2, 100},
			idsToFind: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 20, 30, 40, 50,
				60, 70, 80, 90, 100},
			correctAnswers: []bool{false, true, true, false, true, false, true,
				true, false, true, true, true, false, true, true, true, false,
				true, false},
		},
	}
	for i := range tests {
		answers := make([]bool, 0, len(tests[i].correctAnswers))
		rbt := NewRBTree()
		for _, x := range tests[i].idsToInsert {
			rbt.InsertNode(x)
		}
		for _, x := range tests[i].idsToDelete {
			rbt.DeleteNode(x)
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
