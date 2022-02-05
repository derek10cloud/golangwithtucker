package fedex

import "fmt"

type FedexSender struct {
	//...
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("fedex sends %s parcel\n", parcel)
}
