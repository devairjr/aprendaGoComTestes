package contexto

import (
	"context"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store foi cancelado")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Errorf("store não foi avisada para cancelar")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Errorf("store foi avisada para cancelar")
// 	}
// }
