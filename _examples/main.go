package main

import (
	"fmt"
	go_microsoftapis "github.com/BRUHItsABunny/go-microsoftapis"
)

func main() {
	client := go_microsoftapis.GetTranslateClientWithGeo("america")

	fmt.Println(client.Translate("bunny", "en", "es"))
	// fmt.Println(client.Languages("compact"))
	// fmt.Println(client.Transliterate("バニー", "ja", "Jpan", "Latn")) // Returns "baney", tried to translate bunny to japanese and then take their symbols for transliteration
	// fmt.Println(client.Detect("bunny")) // unauthorized, i guess they did protect the SDK from being used by the app when the app doesnt naturally use it
	// fmt.Println(client.DictionaryLookup("Rabbit", "en", "nl"))
	// fmt.Println(client.DictionaryExamples("fly", "volar", "en", "es")) // This doesnt get used by the app naturally yet this works???, unlike detect
	// fmt.Println(client.Speak("Rabbits like to fuck", &go_microsoftapis.SpeakVoice{Language: "en", Voice: "en-CA-Linda", Gender: "female"}))

}
