package generator

import "fmt"

// InMemoryAccountNumberGenerator generator w pamieci
type InMemoryAccountNumberGenerator struct {
	counter int
}

func (generator *InMemoryAccountNumberGenerator) Next() string {
	generator.counter++
	return fmt.Sprintf("%026d", generator.counter)
}
