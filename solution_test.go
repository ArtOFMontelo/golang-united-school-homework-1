package main

import (
	"fmt"
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) string {
	t.Error("Result should be Hello:world_map:")
	return ""
	}	
func TestMain(t *testing.T) {
	helloWorld := emoji.Sprint("Hello :world_map:!")
	fmt.Println(helloWorld)
	t.Error("Result should be Hello:world_map:")
}
