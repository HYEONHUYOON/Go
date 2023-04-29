package koreapost

import "fmt"

type PostSender struct {
}

func (f *PostSender) Sender(parcel string) {
	fmt.Println(parcel)
}
