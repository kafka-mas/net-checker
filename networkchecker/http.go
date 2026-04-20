package networkchecker

// Stub
type HTTPChecker struct{}

func (HTTPChecker) Ping(address []string) ([]string, error) {
	return []string{"Hello from Http"}, nil
}
