package main

import (
	"fmt"
	"math"
	"time"
)

type ErrorNegativeSqrt struct {
	When time.Time
	What string
}

func (e *ErrorNegativeSqrt) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, &ErrorNegativeSqrt{
			time.Now(),
			"error negative sqrt",
		}
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
