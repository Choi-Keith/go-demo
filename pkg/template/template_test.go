package template

import (
	"fmt"
	"testing"
)

func TestGenerateTemplate(t *testing.T) {
	result := generateTemplate()
	fmt.Printf("result: %s", result)
}
