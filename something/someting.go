package something

import "fmt"

// export 안됨
func sayHello() {
	fmt.Println("Hi")
}

// export 됨
func SayHello() {
	fmt.Println("Hi")
}