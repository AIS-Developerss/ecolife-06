# Быстрый старт

## Настройка переменных окружения

Перед первым запуском необходимо настроить переменные окружения для backend:

1. **Backend**: 
```bash
cp backend/.env.example backend/.env
# Отредактируйте backend/.env при необходимости
```

**Примечание:** Frontend - это статический сайт (HTML/CSS/JS), переменные окружения не требуются.

## Запуск через Docker Compose (рекомендуется)

1. Убедитесь, что Docker и Docker Compose установлены

2. Настройте переменные окружения для backend (см. выше)

3. Запустите все сервисы:
```bash
docker-compose up -d
```

4. Дождитесь запуска всех контейнеров (около 30-60 секунд)

5. Откройте в браузере:
   - **Frontend**: http://localhost:8000
   - **Backend API**: http://localhost:8080/api

6. Для просмотра логов:
```bash
docker-compose logs -f
```

7. Для остановки:
```bash
docker-compose down
```

## Локальная разработка

### Backend

1. Установите PostgreSQL локально или используйте Docker:
```bash
docker-compose up -d postgres
```

2. Создайте файл `backend/.env` (скопируйте из `.env.example`)

3. Запустите миграции:
```bash
cd backend
make migrate-up
# или вручную:
psql -U postgres -d ecolife -f migrations/001_init_schema.up.sql
```

4. Установите зависимости и запустите:
```bash
go mod download
go run cmd/api/main.go
```

Сервер будет доступен на http://localhost:8080

### Frontend

Фронтенд - это статический сайт, можно открыть напрямую в браузере:

1. Откройте `frontend/index.html` в браузере

2. Или используйте простой HTTP сервер:
```bash
cd frontend
python3 -m http.server 8000
```

3. Откройте http://localhost:8000

**Важно**: При локальной разработке фронтенда убедитесь, что в `frontend/index.html` указан правильный URL API в атрибуте `data-action` формы (по умолчанию `http://localhost:8080/api/feedback`).

## Проверка работы API

```bash
# Создать заявку из формы обратной связи
curl -X POST http://localhost:8080/api/feedback \
  -H "Content-Type: application/json" \
  -H "Origin: http://localhost:8000" \
  -d '{
    "name": "Иван Иванов",
    "phone": "+79991234567"
  }'

# Проверить созданные заявки в БД
docker-compose exec postgres psql -U postgres -d ecolife -c "SELECT * FROM applications ORDER BY created_at DESC LIMIT 5;"
```

## Структура базы данных

После запуска миграций будет создана таблица:
- `applications` - заявки на услуги из формы обратной связи

Миграция применяется автоматически при первом запуске PostgreSQL через Docker Compose.

## Полезные команды

```bash
# Просмотр логов всех сервисов
docker-compose logs -f

# Просмотр логов конкретного сервиса
docker-compose logs -f backend
docker-compose logs -f frontend
docker-compose logs -f postgres

# Перезапуск сервиса
docker-compose restart backend

# Пересборка и перезапуск
docker-compose up -d --build

# Остановка всех сервисов
docker-compose down

# Остановка и удаление volumes (удалит данные БД)
docker-compose down -v
```
