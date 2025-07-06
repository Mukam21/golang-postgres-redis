# Golang Gin PostgreSQL Redis API

Простой REST API на Go с использованием Gin, PostgreSQL и Redis.

---

## Стек технологий

- Go 1.24

- Gin (HTTP-фреймворк)

- PostgreSQL (БД)

- Redis (Кэш)

- GORM (ORM)

- Docker

---

## Быстрый старт

### 1. Создайте файл `.env` в корне проекта с такими значениями:

```env

            DB_HOST=...

            DB_USER=...

            PASSWORD=...

            DBNAME=...

            PORT=...

            REDIS_ADDR=...

2. Соберите Docker-образ:

```bash
          docker build -t golang_redis .

3. Запустите контейнер:

          docker run --env-file .env -p 8080:8080 golang_redis

API будет доступен по адресу: http://localhost:8080

Проверка API через Postman:

1. Создать пользователя (POST)

Метод: POST

URL: http://localhost:8080/users?name=Alice&email=alice@example.com

Описание: Создаёт нового пользователя с параметрами name и email в query строке.

Ответ:

                        {
                           "ID": 1,
                           "Name": "Alice",
                           "Email": "alice@example.com"
                        }


2. Получить пользователя по ID (GET)

Метод: GET

URL: http://localhost:8080/users/1

Описание: Получает пользователя с ID=1 (замените на нужный ID)

Ответ (из базы или кэша):

                        From DB: Alice <alice@example.com>

или

                        From Cache: Alice <alice@example.com>
