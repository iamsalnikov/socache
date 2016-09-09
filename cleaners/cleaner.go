package cleaners

// Cleaner must be implemented by cleaners
type Cleaner interface {
	Clear(url string) (bool, error)
}
