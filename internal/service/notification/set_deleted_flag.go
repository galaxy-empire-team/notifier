package notification

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) SetDeletedFlag(ctx context.Context, userID uuid.UUID, notificationIDs []uint64) error {
	err := s.notificationStorage.SetDeletedFlag(ctx, userID, notificationIDs)
	if err != nil {
		return fmt.Errorf("notificationStorage.SetDeletedFlag(): %w", err)
	}

	return nil
}
