# Concurrency Parser

Параллельный парсер новостей. Одновременно парсит несколько сайтов, сохраняет результаты в PostgreSQL, отдаёт через REST API.

## Технологии

- Go
- PostgreSQL
- `net/http`
- `github.com/gocolly/colly`
- `github.com/lib/pq`

## Запуск

### 1. Требования

- Go 1.21+
- PostgreSQL

### 2. Создай базу данных

```sql
CREATE DATABASE responses;
```

### 3. Настрой переменную окружения

```powershell
# Windows PowerShell
$env:DB_PASSWORD="твойпароль"
```

```bash
# Linux / macOS
export DB_PASSWORD=твойпароль
```

### 4. Установи зависимости

```bash
go mod tidy
```

### 5. Запусти

```bash
go run .
```

При запуске программа параллельно парсит ria.ru и rbc.ru, сохраняет результаты в БД. Сервер запускается на `http://localhost:8080`.

---

## Эндпоинты

| Метод | URL | Описание |
|---|---|---|
| GET | `/result` | Получить все новости |

### Пример ответа

```json
[
  {
    "id": 1,
    "header": "Заголовок новости",
    "link": "https://ria.ru/..."
  }
]
```

---

## Структура проекта

```
concurrencyParser/
├── main.go
├── models/
│   └── models.go
├── storage/
│   └── storage.go
├── parser/
│   └── parser.go
└── handlers/
    └── handlers.go
```
