package main

import (
	"github.com/VelHRH/data-structures/tree/main/linkedlist"
)

func main() {
	list := linkedlist.LinkedList{}
	list.Prepend(30)
	list.Prepend(40)
	list.Prepend(15)
	list.Prepend(4)
	list.Prepend(22)
	list.Print()
	list.DeleteByValue(41)
	list.Print()
}
