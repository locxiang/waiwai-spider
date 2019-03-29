package model

func MigrateInit() error {

	var models = []IModel{
		&Book{},
		&Chapter{},
		&Detail{},
	}

	for _, m := range models {
		err := m.Migrate()
		if err != nil {
			return err
		}
	}

	return nil
}
