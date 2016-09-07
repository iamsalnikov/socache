package cleaners

/// CleanerInterface must be implemented by cleaners
type CleanerInterface interface {
	Clear(url string) (bool, error)
}
