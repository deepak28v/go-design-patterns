package singleton

// Counter interface
type Counter interface {
	AddOne() int
}

type counter struct {
	count int
}

var instance *counter

// GetInstance to return a Singleton
func GetInstance() Counter {
	if instance == nil {
		//instance = &singleton{}
		instance = new(counter)
	}
	return instance
}

func (s *counter) AddOne() int {
	s.count++
	return s.count
}
