package main
import(
	
	"regexp"
)
func punc(s string) string {
	re := regexp.MustCompile(`\s*([.,;:?!]+)`)
	s = re.ReplaceAllString( s,"$1")
	re1 := regexp.MustCompile(`([,.?!;:]+)(\w)`)
	s = re1.ReplaceAllString(s, (`$1 $2`))
	return s
}
