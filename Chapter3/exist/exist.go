package exist

func Exist(words []string, s string) bool {
    n := len(words)
    l, r := 0, n - 1
    for l < r {
        m := (l + r) / 2
        if words[m] == s {
            return true
        } else if words[m][0] >= s[0] {
            r = m
        } else {
            l = m + 1
        }
    }

    return false
}
