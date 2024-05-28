package builder

import (
	"fmt"
	"strings"
)

func BuildMultipleItemInsertQuery(numItems int, valuesByItem int) string {
	if numItems == 0 || valuesByItem == 0 {
		return ""
	}
	idx := 0
	groups := make([]string, 0, numItems)
	for i := 0; i < numItems; i++ {
		ids := make([]any, valuesByItem)
		placeholders := make([]string, valuesByItem)

		for j := 0; j < valuesByItem; j++ {
			idx++
			ids[j] = idx
			placeholders[j] = "$%d"
		}
		formattedQuery := "(" + strings.Join(placeholders, ", ") + ")"
		groups = append(groups, fmt.Sprintf(formattedQuery, ids...))
	}

	return strings.Join(groups, ", ")
}
