package leetcode

import (
	"fmt"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	five := IntListNode{
		Val:  5,
		Next: nil,
	}
	four := IntListNode{
		Val:  4,
		Next: &five,
	}
	three := IntListNode{
		Val:  3,
		Next: &four,
	}
	two := IntListNode{
		Val:  2,
		Next: &three,
	}
	one := IntListNode{
		Val:  1,
		Next: &two,
	}

	reversed := reverseKGroup(&one, 3)
	fmt.Println(reversed)
}
