package main

import (
	"rainmock/mock"
)

func main() {
	go mock.Watch()
	mock.StartCommand()
}
