package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAll(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting all data: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
