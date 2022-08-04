package pike

// Sorted is to help split out permission to the relevant auth
type Sorted struct {
	AWS   []string
	GCP   []string
	AZURE []string
}

// ResourceV2 is what resources get parsed into
type ResourceV2 struct {
	TypeName     string
	Name         string
	ResourceName string
	Provider     string
	Attributes   []string
}
