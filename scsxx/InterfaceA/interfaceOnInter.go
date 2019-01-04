package InterfaceA

import "fmt"

type InterfaceOnInter interface {
	InterfaceOnInterA
}

type InterfaceOnInterA interface {
	QueryItemName()
}


type InterfaceOnIn struct {
}

func (this *InterfaceOnIn) QueryItemName() {
		fmt.Println("进入QueryItemName")
}


