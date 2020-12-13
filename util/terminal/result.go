package terminal

type Result struct {
	Origin  string
	Method  string
	Success bool
}

func NewResult(origin string, method string, success bool) *Result {
	return &Result{origin, method, success}
}
