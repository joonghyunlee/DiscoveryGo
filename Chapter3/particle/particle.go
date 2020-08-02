package particle

import (
    "unicode/utf8"
)

var (
    start = rune(44032)
)

func HasFinalConsonant(s string) bool {
    numEnds := 28
    r, _ := utf8.DecodeLastRuneInString(s)
    return int(r - start) % numEnds != 0
}
