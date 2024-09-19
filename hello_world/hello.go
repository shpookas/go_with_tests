package main

const spanish = "Spanish"
const french = "French"
const lithuanian = "Lithuanian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const lithuanianHelloPrefix = "Labas, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case lithuanian:
		prefix = lithuanianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
