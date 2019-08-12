package hello

import (
	"fmt"
)

func Hello() {
	fmt.Println("Hello, world!")
}

func CustomHello(name string) {
	fmt.Printf("Hello, %s!", name)
}
