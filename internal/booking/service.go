package booking

import "context"

type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
	return &Service{store}
}

func (s *Service) Book(b Booking) (Booking, error) {
	session, err := s.store.Book(b)
	if err != nil {
		return Booking{}, err
	}

	return session, nil
}

func (s *Service) ListBookings(movieID string) ([]Booking, error) {

	bookings, err := s.store.ListBookings(movieID)

	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *Service) ConfirmSeat(ctx context.Context, sessionID string, userID string) (Booking, error) {
	return s.store.Confirm(ctx, sessionID, userID)
}

func (s *Service) ReleaseSeat(ctx context.Context, sessionID string, userID string) error {
	return s.store.Release(ctx, sessionID, userID)
}
