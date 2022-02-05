package koreaPost

import "fmt"

type PostSender struct {
	//..
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %s를 보냅니다.\n", parcel)
}
