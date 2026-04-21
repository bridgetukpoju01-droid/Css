package main
import(

	"regexp"
)
func fixquotes(s string) string {
	w := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = w.ReplaceAllString(s, "'$1'")
	return s
}
