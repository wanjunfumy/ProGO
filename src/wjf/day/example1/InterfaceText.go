package main

import "fmt"

type Book interface {
	getName() string
	isRead() bool
}

type Xiyouji struct {
	Name string
}

func (x Xiyouji) getName() string {
	return x.Name
}

func (x Xiyouji) isRead() bool {
	return true
}

func main() {
	var xi Xiyouji
	xi.Name = "四大名著之一"
	fmt.Println(xi.getName())
	fmt.Println(xi.isRead())
}
