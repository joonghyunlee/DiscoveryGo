package calc

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// BinOp type is a binary operation.
type BinOp func(int, int) int

// StrSet is String set data structure.
type StrSet map[string]struct{}

// PrecMap is a Operation - Higher Priority Operations Map.
type PrecMap map[string]StrSet

// NewStrSet returns a new StrSet instance.
func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

// Eval returns an integer value corresponding to the given expression.
func Eval(opMap map[string]BinOp, prec PrecMap, expr string) (int, error) {
	ops := []string{"("}
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}

			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}

	for _, token := range strings.Fields(expr) {
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			reduce(token)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, errors.New("Invalid Parameter")
			}
			nums = append(nums, num)
		}
	}

	reduce(")")
	return nums[0], nil
}

// NewEvaluator returns the Eval function.
func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) (int, error) {
	return func(expr string) (int, error) {
		return Eval(opMap, prec, expr)
	}
}

// ReplaceAll returns the given string.
func ReplaceAll(in string) string {
	rx := regexp.MustCompile(`{[^}]+}`)
	out := rx.ReplaceAllStringFunc(in, func(expr string) string {
		expr = strings.Trim(expr, "{ }")
		eval := NewEvaluator(map[string]BinOp{
			"**": func(a, b int) int {
				if a == 1 {
					return 1
				}
				if b < 0 {
					return 0
				}
				r := 1
				for i := 0; i < b; i++ {
					r *= a
				}
				return r
			},
			"*":   func(a, b int) int { return a * b },
			"/":   func(a, b int) int { return a / b },
			"mod": func(a, b int) int { return a % b },
			"+":   func(a, b int) int { return a + b },
			"-":   func(a, b int) int { return a - b },
		}, PrecMap{
			"**":  NewStrSet(),
			"*":   NewStrSet("**", "*", "/", "mod"),
			"/":   NewStrSet("**", "*", "/", "mod"),
			"mod": NewStrSet("**", "*", "/", "mod"),
			"+":   NewStrSet("**", "*", "/", "mod", "+", "-"),
			"-":   NewStrSet("**", "*", "/", "mod", "+", "-"),
		})

		r, err := eval(expr)
		if err != nil {
			return err.Error()
		}

		return strconv.Itoa(r)
	})

	return out
}
