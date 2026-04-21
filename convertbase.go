package main
import(
	
	"strings"
	"strconv"
)
func convertbase(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++{
		switch strings.ToLower(words[i]) {
	case "(hex)":
		if dec,err := strconv.ParseInt(words[i-1],16,64); err == nil {
			words[i-1] = strconv.FormatInt(dec,10)
			words = append(words[:i],words[i+1:]...)
		}
	case "(bin)":
		if dec,err := strconv.ParseInt(words[i-1],2,64); err == nil {
			words[i-1] = strconv.FormatInt(dec,10)
			words = append(words[:i],words[i+1:]...)
		}
	}
	
}
return strings.Join(words," ")
	} 
	