# Тестовое задание — API для управления задачами и пользователями

## Описание

Это REST API для управления пользователями и задачами с авторизацией и системой очков.
Основные функции: регистрация, авторизация, добавление и просмотр задач, отметка задач как выполненных, просмотр статуса пользователя и таблицы лидеров, установка реферера.

---

## Установка
1.	Клонируйте репозиторий:
```bash
git clone https://github.com/labelq/test-task.git
cd test-task
```
2.	Убедитесь, что у вас установлен и запущен PostgreSQL (настройки подключения в .env).

3. Запустите сервер с использованием Docker:
```bash
docker-compose up --build
```

**ВАЖНО!**  
После первого запуска контейнеров необходимо применить миграции для создания таблиц в базе данных. Выполните:
```bash
make migrate-up
```
Если у вас нет утилиты migrate, установите её согласно [инструкции](https://github.com/golang-migrate/migrate).

Теперь API готов к работе!

---

## Работа с API

- Все запросы выполняются по адресу http://localhost:8000.
- Все запросы кроме sign-up и sign-in требуют авторизации через Bearer Token с использованием токена, 
который выдается при авторизации.

### 1. Регистрация пользователя
- Метод: POST
- URL: /auth/sign-up
- Тело запроса (raw JSON):
```bash
{
    "name": "Alex",
    "password": "qwerty"
}
```
Ответ:
Вернётся JSON с id зарегистрированного пользователя.
---

### 2. Авторизация пользователя

- Метод: POST
- URL: /auth/sign-in
- Тело запроса (raw JSON):
```bash
{
"name": "Alex",
"password": "qwerty"
}
```

Ответ:
Токен авторизации (Bearer token), который нужно использовать в заголовке Authorization для следующих запросов.

---

### 3. Добавление задачи

- Метод: POST
- URL: /api/tasks/add
- Authorization: Bearer <your_token>

Тело запроса (raw JSON):
```bash
{
"name": "Telegram",
"instruction": "Subscribe @ssorval"
}
```
Ответ:
Вернётся id созданной задачи.
---

### 4. Просмотр всех задач

- Метод: GET
- URL: /api/tasks/list
- Authorization: Bearer <your_token>

Ответ:
Список всех задач с их описаниями и очками.

---

### 5. Получить информацию о пользователе

- Метод: GET
- URL: /api/users/:id/status
- Authorization: Bearer <your_token>
- Пример: /api/users/1/status

Ответ:
Статус пользователя, включая количество очков и прочее.

---

### 6. Таблица лидеров

- Метод: GET
- URL: /api/users/leaderboard
- Authorization: Bearer <your_token>

Ответ:
Топ-3 пользователя по количеству очков.

---

### 7. Отметить задачу как выполненную

- Метод: POST
- URL: /api/users/:id/task/complete
- Authorization: Bearer <your_token>
- Пример: /api/users/1/task/complete

Тело запроса (raw JSON):
```bash
{
"task_id": 1,
"complete": true
}
```
Ответ:
Обновлённый статус пользователя и очки.
---

### 8. Добавить реферера пользователю

- Метод: POST
- URL: /api/users/:id/referrer
- Authorization: Bearer <your_token>
- Пример: /api/users/2/referrer

Тело запроса (raw JSON):
```bash
{
"referrer": 1
}
```

Ответ:
Обновлённая информация о пользователе с привязанным реферером.

