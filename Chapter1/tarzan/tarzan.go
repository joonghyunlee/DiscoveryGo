package main

import "fmt"

func song(n int) {
	priceBrief := 10
	priceKnife := 20

	for i := 0; i < n; i++ {
		message := "타잔이 %d원짜리 팬티를 입고, "
		message += "%d원짜리 칼을 차고 노래를 한다.\n"
		priceBrief += 10
		priceKnife += 10
		fmt.Printf(message, priceBrief, priceKnife)
	}
}

func main() {
	song(10)
}
