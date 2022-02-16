package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	wM := emoji.Sprint("Hello :world_map:!")
	return wM
}
