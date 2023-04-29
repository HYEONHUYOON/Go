package main

import (
	"interface/fedex"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

type PostSend interface {
	Send(parcel string)
}

func main() {
	var sender PostSend = &fedex.FedexSender{}
	sender.Send("어린왕자")
}
