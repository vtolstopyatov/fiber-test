# Базовый образ
FROM golang

# Установка Goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Создание директории для миграций
WORKDIR /app/migrations

# Копирование файлов миграций
COPY ./migrations .

# Команда для применения миграций
CMD goose -dir . postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${DB_NAME} host=${DB_HOST} port=${DB_PORT} sslmode=disable" up