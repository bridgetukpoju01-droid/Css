package main
import(
	
	"strings"
)
func vowel(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++{
		v := "aeiouhAEIOUH"
		n := words[i+1]
		switch words[i]{
			case"A":
		if strings.ContainsAny(string(n[0]),v) {
			words[i] = "An"
		}
	case"a":
		if strings.ContainsAny(string(n[0]),v) {
			words[i]= "an"
		}
		}
	
	}
	return strings.Join(words, " ")
}

