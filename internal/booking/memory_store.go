package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: make(map[string]Booking),
	}
}

func (s *MemoryStore) Book(b Booking) (Booking, error) {

	_, ok := s.bookings[b.SeatID]
	if ok {
		return Booking{}, ErrSeatAlredyBooked
	}

	s.bookings[b.SeatID] = b

	return b, nil

}
func (s *MemoryStore) ListBookings(movieID string) []Booking {

	out := make([]Booking, 0)

	for _, booking := range s.bookings {
		if booking.MovieID == movieID {
			out = append(out, booking)
		}
	}

	return out

}
