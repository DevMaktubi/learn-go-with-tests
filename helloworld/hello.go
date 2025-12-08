package hello

import "fmt"

const prefixEnglish = "Hello, %s!"
const prefixSpanish = "Hola, %s!"
const prefixFrench = "Bonjour, %s!"

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := prefixEnglish
	switch lang {
		case "Spanish":
			prefix = prefixSpanish
		case "French":
			prefix = prefixFrench
	}
	return fmt.Sprintf(prefix,name)
}