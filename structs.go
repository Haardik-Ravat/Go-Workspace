package main

import "fmt"

type car struct {
	gaspedal uint16
	brake    uint16
	wheel    int16
	topspeed float64
}

func main() {
	acar := car{gaspedal: 2234,
		brake:    3453,
		wheel:    345,
		topspeed: 345345}

	fmt.Println(acar.wheel)

}
