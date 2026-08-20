package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask14(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	in := Inspection{ID: "i", StoreID: "s", CompletedAt: now.Add(-5 * 24 * time.Hour)}
	a := Appeal{InspectionID: "i", StoreID: "s", Reason: "new evidence", Evidence: []string{"photo"}}
	require.NoError(t, s.CheckAppeal(context.Background(), a, in))
}
