package queue

import "fmt"

// Eval returns characters from the given string.
func Eval(s string) {
	var queue []rune

	push := func(c rune) {
		queue = append(queue, c)
	}

	pop := func() rune {
		c := queue[0]
		queue = queue[1:]

		return c
	}

	isEmpty := func() bool {
		return len(queue) == 0
	}

	for _, c := range s {
		push(c)
	}

	for !isEmpty() {
		fmt.Printf("%c\n", pop())
	}
}
