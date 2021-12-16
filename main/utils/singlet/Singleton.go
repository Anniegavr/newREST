package singlet

import "sync"

type hash map[string]interface{}

type singleton struct {
	items hash
	mutex sync.RWMutex
}

func (s *singleton) Set(key string, data interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.items[key] = data
}

func (s *singleton) Get(key string) (interface{}, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	item, ok := s.items[key]
	if !ok {
		return nil, false
	}

	return item, true
}

var (
	container *singleton
	once      sync.Once
)

func Singleton() *singleton {
	once.Do(func() {
		if container == nil {
			container = &singleton{
				items: make(hash),
			}
		}
	})

	return container
}
