package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return getCorrectLanguagePrefix(language) + name
}

func getCorrectLanguagePrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Anagh", ""))
}
