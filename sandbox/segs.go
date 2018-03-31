package main

import (
	"fmt"
)

type InterRelate uint8

const (
	OtherB InterRelate = 1 << iota // 1 << 0 == 0001
	OtherA                         // 1 << 1 == 0010
	SelfB                          // 1 << 2 == 0100
	SelfA                          // 1 << 3 == 1000
)
const (
	SelfMask  = SelfA | SelfB
	OtherMask = OtherA | OtherB
)

func isSelfPresent(x InterRelate) {
	fmt.Printf("%04x\n", x)
	fmt.Printf("%04x\n", SelfMask)
	fmt.Printf("%04x\n", x&SelfMask)
	fmt.Printf("%d\n", x&SelfMask)
}

func main() {
	fmt.Println(SelfMask)
	var v = SelfA | SelfB | OtherA | OtherB
	fmt.Println(v)
	fmt.Printf("%04b\n", v)
	//fmt.Printf("%04x\n", x)
	//fmt.Println(bits.OnesCount32(x))
	//bs := [5]byte{1, 1, 1, 1}
	//fmt.Println(bs)
}
