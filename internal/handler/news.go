package handler

import (
	"fibertest/internal/db"
	"fibertest/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetNewsList(c *fiber.Ctx) error {
	query := `
	SELECT news.id, news.title, news.content, array_agg(news_categories.category_id) AS categories
	FROM news
	LEFT JOIN news_categories ON news.id = news_categories.news_id
	GROUP BY news.id
	ORDER BY news.id DESC
	LIMIT $1 OFFSET $2
	`

	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	defer rows.Close()

	newsList := []models.News{}
	for rows.Next() {
		news := models.News{}
		rows.Scan(&news.ID, &news.Title, &news.Content, &news.Categories)
		if len(news.Categories) == 0 {
			news.Categories = []int64{}
		}
		newsList = append(newsList, news)
	}

	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"News":    newsList,
	})
}

func GetNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that :id is an integer")
	}

	query := `
	SELECT news.id, news.title, news.content, array_agg(news_categories.category_id) AS categories
	FROM news
	LEFT JOIN news_categories ON news.id = news_categories.news_id
	WHERE news.id = $1
	GROUP BY news.id
	`
	row := db.DB.QueryRow(query, id)

	var news models.News
	err = row.Scan(&news.ID, &news.Title, &news.Content, &news.Categories)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).JSON(news)
}

type NewsUpdate struct {
	ID         *int64  `json:"id"`
	Title      *string `json:"title"`
	Content    *string `json:"content"`
	Categories []int64 `json:"categories"`
}

func UpdateNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please ensure that :id is an integer")
	}

	var newsUpdate NewsUpdate
	if err := c.BodyParser(&newsUpdate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON format")
	}

	if newsUpdate.ID != nil && int(*newsUpdate.ID) != id {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot change news ID")
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	newsDB, err := tx.FindByPrimaryKeyFrom(models.NewsTable, id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "News not found")
	}

	if err := c.BodyParser(&newsDB); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot change news ID")
	}

	err = tx.Update(newsDB)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update news")
	}

	categories := newsUpdate.Categories
	if categories != nil {
		_, err = tx.DeleteFrom(models.NewsCategoriesTable, "WHERE news_id = $1", id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to update news categories")
		}

		for _, category := range categories {
			newsCategories := &models.NewsCategories{
				NewsId:     int64(id),
				CategoryId: category,
			}
			err = tx.Insert(newsCategories)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Failed to update news categories")
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update news")
	}

	return GetNews(c)
}
