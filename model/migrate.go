package model

func MigrateInit() error {
	err := new(Book).Migrate()
	if err != nil {
		return err
	}
	return nil
}
