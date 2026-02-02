# Настройка переменных окружения

## Backend

1. Скопируйте файл `.env.example` в `.env`:
```bash
cp backend/.env.example backend/.env
```

2. Отредактируйте `backend/.env` под ваши настройки:
```env
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecolife
DB_SSLMODE=disable
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
CORS_ALLOWED_ORIGINS=http://localhost:3000
```

**Важно:** Для поддержки нескольких origins в CORS, укажите их через запятую:
```env
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001,https://example.com
```

## Frontend

1. Скопируйте файл `.env.local.example` в `.env.local`:
```bash
cp frontend/.env.local.example frontend/.env.local
```

2. Отредактируйте `frontend/.env.local`:
```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

**Важно:** 
- В Next.js переменные окружения, доступные в браузере, должны начинаться с `NEXT_PUBLIC_`
- Файл `.env.local` не должен попадать в git (уже добавлен в .gitignore)

## Docker Compose

При использовании `docker-compose up`, переменные окружения будут загружены из:
- `backend/.env` для backend сервиса
- `frontend/.env.local` для frontend сервиса

Если файлы не существуют, будут использованы значения по умолчанию из `docker-compose.yml`.

## Переменные окружения

### Backend

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `SERVER_PORT` | Порт для HTTP сервера | `8080` |
| `SERVER_HOST` | Хост для HTTP сервера | `0.0.0.0` |
| `DB_HOST` | Хост PostgreSQL | `localhost` |
| `DB_PORT` | Порт PostgreSQL | `5432` |
| `DB_USER` | Пользователь PostgreSQL | `postgres` |
| `DB_PASSWORD` | Пароль PostgreSQL | `postgres` |
| `DB_NAME` | Имя базы данных | `ecolife` |
| `DB_SSLMODE` | Режим SSL для PostgreSQL | `disable` |
| `REDIS_HOST` | Хост Redis | `localhost` |
| `REDIS_PORT` | Порт Redis | `6379` |
| `REDIS_PASSWORD` | Пароль Redis (пусто для отсутствия) | `` |
| `REDIS_DB` | Номер базы данных Redis | `0` |
| `CORS_ALLOWED_ORIGINS` | Разрешенные origins для CORS (через запятую) | `http://localhost:3000` |
| `ENV` | Окружение (development/production) | `development` |

### Frontend

| Переменная | Описание | По умолчанию |
|------------|----------|--------------|
| `NEXT_PUBLIC_API_URL` | URL бэкенд API | `http://localhost:8080/api` |
| `NODE_ENV` | Окружение Node.js | `development` |

## Безопасность

⚠️ **Важно:** Никогда не коммитьте файлы `.env` и `.env.local` в git! Они уже добавлены в `.gitignore`.

Используйте `.env.example` и `.env.local.example` как шаблоны для других разработчиков.

