package utils

import (
	"testing"
	"fmt"
)

func TestMd5(t *testing.T) {
	str:=Md5("zhaohu")
	fmt.Println(str)
}
