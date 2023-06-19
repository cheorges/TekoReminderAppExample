package main

import "reminders/pkg/controller"

func main() {
	handler := controller.New()
	handler.Run()
}
