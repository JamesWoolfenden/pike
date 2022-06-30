package pike

type template struct {
	Resource string `json:"resource"`
	API      string `json:"api"`
	Filename string `json:"filename"`
	Provider string `json:"provider"`
	Code     string `json:"code"`
}
