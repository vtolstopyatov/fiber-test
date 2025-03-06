-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS news_categories (
    news_id bigint NOT NULL,
    category_id bigint NOT null,
PRIMARY KEY (news_id, category_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS news_categories;
-- +goose StatementEnd
