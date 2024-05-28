package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func WrapIfErr(err error, args ...string) error {
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("db has empty data: %w", err)
	}

	if err != nil {
		str := strings.Join(args, ".")
		return fmt.Errorf("%w. %s", err, str)
	}

	return nil
}
