package main

import (
	"fmt"
	"time"
)

func main() {
	//default switch
	i := 3

	switch i {
	case 2:
		{
			fmt.Println("two")
		}
	default:
		fmt.Println("not two")
	}

	//conditional switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	//type switch
	functionVariable := func(params interface{}) {
		switch params.(type) {
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Println("other")
		}
	}

	functionVariable(6)
	functionVariable("hello")
	functionVariable(6.4)
}
