package auth

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	fmt.Println(GenerateToken("hello", "hominsu"))
}
