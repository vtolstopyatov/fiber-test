package seed_test_data

import (
	"fibertest/internal/db"
	"fibertest/internal/models"
	"fmt"
)

func SeedData() error {
	var count int
    err := db.DB.QueryRow("SELECT COUNT(*) FROM news").Scan(&count)
    if err != nil {
        return err
    }

    if count == 0 {
		for i := 1; i <= 20; i++ {
			news := &models.News{
				Title: fmt.Sprintf("Test title value %v", i),
				Content: fmt.Sprintf("Test content value %v", i),
			}
			err = db.DB.Insert(news)
			if err != nil {
				return err
			}
		}
    }
    return nil
}