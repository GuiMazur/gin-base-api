package models

func ToMigrate() []interface{} {
	return []interface{}{
		&User{},
	}
}
