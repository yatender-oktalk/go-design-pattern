package creational

// returns the single instance of the counter

type singleton struct {
	count int
}

var instance *singleton

// GetInstance singleton pattern
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
