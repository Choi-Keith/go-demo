package util

import "database/sql"

func SqlNullString(str string) sql.NullString {
	return sql.NullString{
		String: str,
		Valid:  len(str) > 0,
	}
}
