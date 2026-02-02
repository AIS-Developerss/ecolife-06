# Ecolife-06 - Чистый мир

Точная копия сайта ООО «Чистый мир» - регионального оператора по обращению с твердыми коммунальными отходами на территории Республики Ингушетия.

## Технологии

### Backend
- **Golang 1.23** - основной язык программирования
- **Gin** - веб-фреймворк
- **PostgreSQL** - база данных
- **Redis** - кэширование
- **Чистая архитектура** - разделение на слои (domain, application, infrastructure, presentation)
- **DDD** - Domain-Driven Design подход

### Frontend
- **Next.js 14** - React фреймворк
- **TypeScript** - типизация
- **Tailwind CSS** - стилизация
- **React Hook Form** - работа с формами

## Структура проекта

```
ecolife-06/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go          # Точка входа приложения
│   ├── internal/
│   │   ├── domain/              # Доменные сущности и интерфейсы
│   │   ├── application/         # Use cases (бизнес-логика)
│   │   ├── infrastructure/      # Реализация репозиториев (PostgreSQL, Redis)
│   │   └── presentation/        # HTTP handlers и middleware
│   ├── migrations/              # SQL миграции
│   └── go.mod
├── frontend/
│   ├── app/                     # Next.js App Router
│   ├── components/              # React компоненты
│   ├── lib/                     # Утилиты и API клиент
│   └── package.json
└── docker-compose.yml
```

## Быстрый старт

### Требования
- Docker и Docker Compose
- Go 1.23+ (для локальной разработки)
- Node.js 20+ (для локальной разработки)

### Настройка переменных окружения

Перед запуском необходимо настроить переменные окружения:

1. **Backend**: Скопируйте `backend/.env.example` в `backend/.env` и настройте под ваше окружение
2. **Frontend**: Скопируйте `frontend/.env.local.example` в `frontend/.env.local` и настройте URL API

Подробная инструкция в файле [ENV_SETUP.md](./ENV_SETUP.md)

### Запуск через Docker Compose

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd ecolife-06
```

2. Настройте переменные окружения (см. выше)

3. Запустите все сервисы:
```bash
docker-compose up -d
```

4. Приложение будет доступно:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - PostgreSQL: localhost:5432
   - Redis: localhost:6379

### Локальная разработка

#### Backend

1. Скопируйте файл `.env.example` в `.env`:
```bash
cp backend/.env.example backend/.env
```

2. Отредактируйте `backend/.env` под ваши настройки (см. `backend/.env.example`)

2. Установите зависимости:
```bash
cd backend
go mod download
```

3. Запустите миграции (используя golang-migrate или напрямую через psql):
```bash
psql -U postgres -d ecolife -f migrations/001_init_schema.up.sql
```

4. Запустите сервер:
```bash
go run cmd/api/main.go
```

#### Frontend

1. Скопируйте файл `.env.local.example` в `.env.local`:
```bash
cp frontend/.env.local.example frontend/.env.local
```

2. Отредактируйте `frontend/.env.local` и укажите URL вашего backend API

3. Установите зависимости:
```bash
cd frontend
npm install
```

4. Запустите dev сервер:
```bash
npm run dev
```

## API Endpoints

### Applications (Заявки)
- `POST /api/applications` - Создать заявку
- `GET /api/applications` - Получить все заявки (с пагинацией)
- `GET /api/applications/:id` - Получить заявку по ID

### Containers (Контейнеры)
- `GET /api/containers` - Получить все контейнеры
- `GET /api/containers/:id` - Получить контейнер по ID

### Benefits (Льготы)
- `GET /api/benefits` - Получить все льготы

### Tariffs (Тарифы)
- `GET /api/tariffs` - Получить все тарифы
- `GET /api/tariffs/current` - Получить текущий тариф

## Особенности реализации

1. **Чистая архитектура**: Разделение на слои обеспечивает независимость бизнес-логики от инфраструктуры
2. **DDD**: Доменные сущности и интерфейсы определены в слое domain
3. **Graceful shutdown**: Корректное завершение работы сервера
4. **CORS**: Настроен для работы с фронтендом
5. **Миграции**: SQL миграции для управления схемой БД
6. **Docker**: Полная контейнеризация для легкого развертывания

## Лицензия

MIT

