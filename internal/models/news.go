//go:generate reform

package models

import "github.com/lib/pq"

//reform:news
type News struct {
	ID         int64         `reform:"id,pk"`
	Title      string        `reform:"title"`
	Content    string        `reform:"content"`
	Categories pq.Int64Array `reform:"-"`
}

//reform:news_categories
type NewsCategories struct {
	NewsId     int64 `reform:"news_id,pk"`
	CategoryId int64 `reform:"category_id"`
}
