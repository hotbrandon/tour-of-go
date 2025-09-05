package main

// https://go.dev/tour/methods/11

import (
	"fmt"
)

type Shape interface {
    Area() float64
}

type Square struct {
    Side float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

func GetArea(s Shape) float64 {
    return s.Area()
}
func main(){
	// In main, you can assign a Square to a Shape interface
	mySquare := Square{Side: 10}
	fmt.Printf("type of mySquare: %T\n", mySquare) // type of mySquare: main.Square
	fmt.Println(GetArea(mySquare)) // GetArea works with a Square
}
