package analyzer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func Test_Something(t *testing.T) {
	var keyword = Symbol

	//fmt.Println(html.EscapeString("{"))
	assert.Equal(t, keyword.String(), "symbol")

	var bytes = `addffb // this is a line comment
this is outside the comments
aaaa
/* this */
/* this
   is
   a
   multi-line
   comment */`

	re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	newBytes := re.ReplaceAllString(bytes, "")
	fmt.Println(newBytes)
}
