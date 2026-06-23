package booking

import (
	"sync"
)

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: make(map[string]Booking),
	}
}

func (s *ConcurrentStore) Book(b Booking) (Booking, error) {

	s.Lock()
	defer s.Unlock()

	_, ok := s.bookings[b.SeatID]
	if ok {
		return Booking{}, ErrSeatAlredyBooked
	}

	s.bookings[b.SeatID] = b

	return b, nil

}
func (s *ConcurrentStore) ListBookings(movieID string) []Booking {

	s.RLock()
	defer s.RUnlock()

	out := make([]Booking, 0)

	for _, booking := range s.bookings {
		if booking.MovieID == movieID {
			out = append(out, booking)
		}
	}

	return out

}
