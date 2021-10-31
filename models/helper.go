package models

func nullableValueOf(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}
