package main

import (
	"fmt"
	"worthifyme/reward"
)

func main() {
	r := reward.New("shoes", "training shoes")
	fmt.Printf("%#v\n", r.Get())
}
