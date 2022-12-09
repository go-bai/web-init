package pkg

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v any) {
	bs, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bs))
}
