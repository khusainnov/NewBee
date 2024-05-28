package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildMultipleItemInsertQuery(t *testing.T) {
	cases := []struct {
		name          string
		numItems      int
		numFields     int
		expectedQuery string
	}{
		{
			name:          "single_item_single_field",
			numItems:      1,
			numFields:     1,
			expectedQuery: `($1)`,
		},
		{
			name:          "multiple_items_multiple_fields",
			numItems:      2,
			numFields:     3,
			expectedQuery: `($1, $2, $3), ($4, $5, $6)`,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			actualQuery := BuildMultipleItemInsertQuery(test.numItems, test.numFields)
			assert.Equal(t, test.expectedQuery, actualQuery)
		})
	}
}
