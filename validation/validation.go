package validation

// Checker validator interface
type Checker interface {
	Check(value interface{}) error
}
