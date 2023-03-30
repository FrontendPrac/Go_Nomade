// 패키지
package main

// 임포트
import (
	"fmt"
	"strings"
)

// 함수 예시
func multiply(a int, b int) int {
	return a+b
}

// 리턴값 2개 함수 예시
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 리턴값 미리 설정하기
func lenAndUpper_2(name string) (lenght int,uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// 인자 여러개 받기 
func repeatMe(words ...string) {
	fmt.Println(words)
}

// 함수가 끝나고 추가적인 동작 실행시키기
func multiply_2(a int, b int) int {
	defer fmt.Println("I'm done")
	return a+b
}

// 루프 : for ~ range
func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	return total
}

// 조건문(if) : 조건문 작성 이전에 해당 스코프에서만 사용하는 변수를 만들 수 있다.
func canIDrink(age int) bool {
	if koreanAge := age+2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}

// 조건문(switch)
func canIDrink_2(age int) bool {
	switch {
	case age < 18:
		return false
	case age >= 18:
		return true
	}
	return false
}

// 메인 함수
func main() {
	totalLenght, upperName := lenAndUpper("kidongg")
	totalLenght_2, upperName_2 := lenAndUpper_2(("kidongg"))
	total := superAdd(1,2,3,4,5,6)

	 //출력
	fmt.Println(multiply(2,2))
	fmt.Println(totalLenght, upperName)
	fmt.Println(totalLenght_2, upperName_2)
	repeatMe("messi", "ronaldo", "pepe")
	fmt.Println(multiply_2(2,2))
	fmt.Println(total)
	fmt.Println(canIDrink(15))
	fmt.Println(canIDrink_2(18))

	// 메모리 접근
	a := 2
	b := &a
	fmt.Println(a, *b)
	fmt.Println(&a)

	// array : 길이 고정 배열
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "alalal"
	names[4] = "blblbl"
	fmt.Println(names)

	// slice : 길이 무제한 배열
	names_2 := []string{"nico", "lynn", "dal"}
	names_2 = append(names_2, "dazy")
	fmt.Println(names_2)

	// map
	kidongg := map[string]string{"name": "kidongg", "age": "28"}
	fmt.Println(kidongg)
	for key, value := range kidongg {
		fmt.Println(key, value)
	}

	// struct 타입
	type person struct {
		name string
		age int
		favFood []string
	}

	// struct
	favFood := []string{"kimchi", "banana"}
	nico := person{name: "kidongg", age: 18, favFood: favFood}
	fmt.Println(nico.name)
}