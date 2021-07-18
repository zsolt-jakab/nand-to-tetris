package analyzer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"html"
	"testing"
)

func Test_Something(t *testing.T) {
	var keyword = Symbol

	fmt.Println(html.EscapeString("{"))
	assert.Equal(t, keyword.String(), "symbol")
}
