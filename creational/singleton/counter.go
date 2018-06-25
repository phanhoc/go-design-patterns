package singleton

type Counter interface {
	AddOne() int
}
type counter struct {
	count int
}

var instance *counter

func GetInstance() Counter {
	if instance == nil {
		instance = new(counter)
	}

	return instance
}

func (s *counter) AddOne() int {
	s.count++
	return s.count
}
