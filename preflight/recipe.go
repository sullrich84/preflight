package preflight

type Recipe struct {
	Target  string
	Origins []string
	Methods []string
	Headers []string
}
