package sort

import "fmt"

func Example_sort() {
    nums := []int{3, 5, 2, 6, 1, 7}
    sort(nums)
    fmt.Println(nums)
    // Output:
    // [1 2 3 5 6 7]
}
