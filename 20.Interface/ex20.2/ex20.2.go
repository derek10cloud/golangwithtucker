package main

import "github.com/derek10cloud/golangwithtucker/20.Interface/fedex"

type Sender interface {
	Send(parcel string) //각 객체에 Send라는 메서드가 있으면, 그 객체를 Sender로 쓸 수 있겠다.
	//fedex와 koreaPost를 Sender 인터페이스로 대신 사용하는 느낌
}

// 각 struct에게 공통된 메서드가 있으면, 인터페이스로 만들 수 있고, 인터페이스를 각 스트럭트처럼 사용하면 된다.

func SendBook(name string, sender Sender) {
	//원래 sender 매개변수의 타입이 *korePost.PostSender 또는 *fedex.FedexSender였는데 Sender인터페이스로 통합되었다.
	sender.Send(name)
}

func main() {
	//sender := &koreaPost.PostSender{} //인스턴스 하나 생성
	var sender Sender = &fedex.FedexSender{}
	SendBook("어린 왕자", sender) //각 인스턴스 인자로 입력하면, sender의 메서드인 Send메서드 이용 가능
	SendBook("그리스인 조르바", sender)
}

// 인터페이스를 활용하면, 서로 다른 패키지의 struct 타입을 다 각 패키지에 맞게 변경할 필요가 없다. (인터페이스 타입이 서로 다른 패키지의 struct를 포함하고 있으니까)
// 실제 패키지 선언부(import)와 각 변수 초기화 부분만 사용하려는 패키지에 맞게 변경하면 된다. (매개 변수 의 타입 같은 것은 인터페이스로 선언되어 있기 때문에 변경할 필요 없음)
