package exist

import "fmt"

func Example_exist() {
    words := []string{"kilo", "liberty", "mitaka", "newton", "ocata", "pike", "queens"}
    r := Exist(words, "rocky")
    fmt.Println(r)
    r = Exist(words, "mitaka")
    fmt.Println(r)
    // Output:
    // false
    // true
}
