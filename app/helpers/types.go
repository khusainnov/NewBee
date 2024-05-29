package helpers

import "fmt"

type Nums interface {
	int16 | int32 | int64 | int |
		float32 | float64
}

func NumsToString[T Nums](v T) string {
	return fmt.Sprintf("%s", v)
}

func GetStringSafely(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func GetPointerString(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}
