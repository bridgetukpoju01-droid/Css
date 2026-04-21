package main

import (
	"regexp"
	"strings"
)

func hello(s string) string {
	w := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = w.ReplaceAllString(s, "'$1'")

	a := regexp.MustCompile(`"\s+(.*?)\s+"`)
	s = a.ReplaceAllString(s, `$1`)

	c := regexp.MustCompile(`\s+([.,;:!?]+)`)
	s = c.ReplaceAllString(s, `$1`)
	b := regexp.MustCompile(`([,.;:!?]+)\w`)
	s = b.ReplaceAllString(s, `$1 $2`)
	return strings.Join(strings.Fields(s), " ")
}



	l := regexp.MustCompile(`'\s+(.*?)\s+'`)
	s = l.ReplaceAllString(s,"'$1'")

	m := regexp.MustCompile(`"\s+(.*?)\s+"`)
	s = m.ReplaceAllString(s,`$1`)

	a := regexp.MustCompile(`\s+([.,;:!?]+)`)
	s = a.ReplaceAllString(s,`$1`)
	p := regexp.MustCompile(`([.,;:!?]+)\w`)
	s = p.ReplaceAllString(s,`$1 $2`)
	return strings.Join(strings.Fields(s), " ")

