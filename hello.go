package main

import "fmt"

const prefix = "Hello, %s!"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf(prefix, name)
}