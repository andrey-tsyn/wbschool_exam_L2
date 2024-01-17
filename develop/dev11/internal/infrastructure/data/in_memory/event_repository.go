package in_memory

import (
	"sync"
	"task11/internal/domain/model"
	"task11/internal/domain/repository"
	"task11/pkg/map_utils"
	"task11/pkg/slice_utils"
	"time"
)

// InMemoryEventRepository is not optimized in_memory realization of event repository,
// recommended for testing only.
type InMemoryEventRepository struct {
	mutex       sync.RWMutex
	events      map[int]model.Event
	availableId int
}

func NewInMemoryEventRepository() *InMemoryEventRepository {
	return &InMemoryEventRepository{
		mutex:       sync.RWMutex{},
		events:      make(map[int]model.Event),
		availableId: 1,
	}
}

func (i *InMemoryEventRepository) Update(event model.Event) error {
	i.mutex.Lock()
	if _, ok := i.events[event.Id]; !ok {
		i.mutex.Unlock()
		return repository.EventNotFoundError{EventId: event.Id}
	}

	updatedEvent := i.events[event.Id]
	updatedEvent.Date = event.Date
	updatedEvent.Name = event.Name

	i.events[event.Id] = updatedEvent
	i.mutex.Unlock()

	return nil
}

func (i *InMemoryEventRepository) Create(event model.Event) (int, error) {
	i.mutex.Lock()
	event.Id = i.availableId
	i.events[i.availableId] = event
	i.availableId++
	i.mutex.Unlock()

	return event.Id, nil
}

func (i *InMemoryEventRepository) Delete(id int) error {
	i.mutex.Lock()
	if _, ok := i.events[id]; !ok {
		i.mutex.Unlock()
		return repository.EventNotFoundError{EventId: id}
	}

	delete(i.events, id)
	i.mutex.Unlock()

	return nil
}

func (i *InMemoryEventRepository) GetByUserId(userId int, from, to time.Time) ([]model.Event, error) {
	i.mutex.RLock()
	values := map_utils.GetKeys(i.events)
	i.mutex.RUnlock()

	userEvents := slice_utils.Filter(values, func(event model.Event) bool {
		return event.UserId == userId && event.Date.Add(1*time.Second).After(from) && event.Date.Add(-1*time.Second).Before(to)
	})

	return userEvents, nil
}
