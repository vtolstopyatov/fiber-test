# fiber-test

### Тестовый проект с 2 ручками:

`POST /news/:id` — изменение новости. Требуется заголовок `Authorization: Bearer my-super-secret-key`

`GET /news?limit=10&offset=0` — получение списка новостей

### В проекте используются:

- Fiber — HTTP-фреймворк
- Reform — ORM
- PostgreSQL — база данных
- Goose — миграции

### Запуск проекта:

1. Склонировать проект в локальную директорию
2. Создать `.env` файл в корне преоекта с переменными окружения по шаблону `.env.example`
3. Запустить Docker Compose: `docker compose up`

При запуске с переменной окружения `ENVIRONMENT=development` таблица news будет заполнена тестовыми данными, если она была пуста
