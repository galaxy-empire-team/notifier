package notification

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) SetReadFlag(ctx context.Context, userID uuid.UUID, notificationIDs []uint64) error {
	err := s.notificationStorage.SetReadFlag(ctx, userID, notificationIDs)
	if err != nil {
		return fmt.Errorf("notificationStorage.SetReadFlag(): %w", err)
	}

	return nil
}
