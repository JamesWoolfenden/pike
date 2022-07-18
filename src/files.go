package pike

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed api.json
var f []byte

func test() {
	var languages []Language
	_ = json.Unmarshal(f, &languages)

	fmt.Println(languages[0].API)
}

// Language struct
type Language struct {
	API      string `json:"api"`
	Resource string `json:"resource"`
}
