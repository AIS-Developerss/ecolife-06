# Быстрый старт

## Настройка переменных окружения

Перед первым запуском необходимо настроить переменные окружения:

1. **Backend**: 
```bash
cp backend/.env.example backend/.env
# Отредактируйте backend/.env при необходимости
```

2. **Frontend**:
```bash
cp frontend/.env.local.example frontend/.env.local
# Отредактируйте frontend/.env.local при необходимости
```

Подробная инструкция: [ENV_SETUP.md](./ENV_SETUP.md)

## Запуск через Docker Compose (рекомендуется)

1. Убедитесь, что Docker и Docker Compose установлены

2. Настройте переменные окружения (см. выше)

3. Запустите все сервисы:
```bash
docker-compose up -d
```

3. Дождитесь запуска всех контейнеров (около 30-60 секунд)

4. Откройте в браузере:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080/api

5. Для просмотра логов:
```bash
docker-compose logs -f
```

6. Для остановки:
```bash
docker-compose down
```

## Локальная разработка

### Backend

1. Установите PostgreSQL и Redis локально или используйте Docker:
```bash
docker-compose up -d postgres redis
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

### Frontend

1. Установите зависимости:
```bash
cd frontend
npm install
```

2. Запустите dev сервер:
```bash
npm run dev
```

3. Откройте http://localhost:3000

## Проверка работы API

```bash
# Получить все контейнеры
curl http://localhost:8080/api/containers

# Получить все льготы
curl http://localhost:8080/api/benefits

# Получить текущий тариф
curl http://localhost:8080/api/tariffs/current

# Создать заявку
curl -X POST http://localhost:8080/api/applications \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Иванов Иван Иванович",
    "phone": "+79281234567",
    "address": "г. Назрань, ул. Тестовая, д. 1",
    "district": "nazranovsky",
    "service_type": "household"
  }'
```

## Структура базы данных

После запуска миграций будут созданы таблицы:
- `applications` - заявки на услуги
- `containers` - контейнеры для мусора
- `benefits` - льготы
- `tariffs` - тарифы на услуги

Начальные данные загружаются автоматически при миграции.

