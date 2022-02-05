package main

import "fmt"

func main() {
	poet1 := "죽는 날까지 하늘을 우러러 \n\t한 점 부끄럼이 없기를,\n잎애세 이는 바람에도\n"
	//큰따옴표에서 여러 줄을 표현하려면 \n을 사용해야 함
	poet2 := `죽는 날까지 하늘을 우러러
	한점 부끄럼이 없기를,
잎에에 이는 바람에도`
	//백쿼트에서는 여러 줄 표현에 특수 문자가 필요 없음
	fmt.Println(poet1)
	fmt.Println(poet2)

	poet3 := "죽는 날까지 하늘을 우러러 \\n\\t"
	//특수문자 그대로 표현하고 싶으면 \(백슬래쉬)를 한 번 더 사용
	fmt.Println(poet3)
}
