# Ecolife-06 - Чистый мир

Точная копия сайта ООО «Чистый мир» - регионального оператора по обращению с твердыми коммунальными отходами на территории Республики Ингушетия.

## Технологии

### Backend
- **Golang 1.23** - основной язык программирования
- **Gin** - веб-фреймворк
- **PostgreSQL** - база данных
- **Чистая архитектура** - разделение на слои (domain, application, infrastructure, presentation)
- **DDD** - Domain-Driven Design подход

### Frontend
- **HTML/CSS/JavaScript** - статический сайт
- **Tailwind CSS** (CDN) - стилизация
- **Vanilla JavaScript** - работа с формами и API

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
│   │   ├── infrastructure/      # Реализация репозиториев (PostgreSQL, logger)
│   │   └── presentation/        # HTTP handlers и middleware
│   ├── migrations/              # SQL миграции
│   └── go.mod
├── frontend/
│   ├── index.html              # Главная страница
│   ├── css/
│   │   └── main.css            # Стили
│   ├── js/
│   │   └── main.js             # JavaScript логика
│   ├── assets/                 # Изображения и другие ресурсы
│   └── Dockerfile              # Docker образ для фронтенда
└── docker-compose.yml
```

## Быстрый старт

### Требования
- Docker и Docker Compose
- Go 1.23+ (для локальной разработки backend)

### Настройка переменных окружения

Перед запуском необходимо настроить переменные окружения для backend:

1. Скопируйте файл `.env.example` в `.env`:
```bash
cp backend/.env.example backend/.env
```

2. Отредактируйте `backend/.env` под ваши настройки (см. `backend/.env.example`)

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

4. Дождитесь запуска всех контейнеров (около 30-60 секунд)

5. Приложение будет доступно:
   - **Frontend**: http://localhost:8000
   - **Backend API**: http://localhost:8080/api
   - **PostgreSQL**: localhost:5433

6. Для просмотра логов:
```bash
docker-compose logs -f
```

7. Для остановки:
```bash
docker-compose down
```

### Локальная разработка

#### Backend

1. Скопируйте файл `.env.example` в `.env`:
```bash
cp backend/.env.example backend/.env
```

2. Отредактируйте `backend/.env` под ваши настройки

3. Установите зависимости:
```bash
cd backend
go mod download
```

4. Запустите миграции:
```bash
psql -U postgres -d ecolife -f migrations/001_init_schema.up.sql
```

5. Запустите сервер:
```bash
go run cmd/api/main.go
```

#### Frontend

Фронтенд - это статический сайт, можно открыть напрямую в браузере:

1. Откройте `frontend/index.html` в браузере

2. Или используйте простой HTTP сервер:
```bash
cd frontend
python3 -m http.server 8000
```

3. Откройте http://localhost:8000

**Важно**: При локальной разработке фронтенда убедитесь, что в `frontend/index.html` указан правильный URL API в атрибуте `data-action` формы (по умолчанию `http://localhost:8080/api/feedback`).

## API Endpoints

### Feedback (Обратная связь)
- `POST /api/feedback` - Создать заявку из формы обратной связи
  - Тело запроса: `{"name": "Имя", "phone": "+79991234567"}`
  - Ответ: `201 Created` с данными созданной заявки

## Особенности реализации

1. **Чистая архитектура**: Разделение на слои обеспечивает независимость бизнес-логики от инфраструктуры
2. **DDD**: Доменные сущности и интерфейсы определены в слое domain
3. **Graceful shutdown**: Корректное завершение работы сервера
4. **CORS**: Настроен для работы с фронтендом
5. **Rate Limiting**: Ограничение частоты запросов (10 запросов в минуту с одного IP)
6. **Валидация данных**: Проверка формата телефона, имени и других полей
7. **Структурированное логирование**: JSON логи с уровнем, временем и контекстом
8. **Обработка ошибок**: Безопасные сообщения об ошибках на русском языке
9. **Миграции**: SQL миграции для управления схемой БД
10. **Docker**: Полная контейнеризация для легкого развертывания

## Безопасность

- Валидация всех входящих данных
- Защита от SQL-инъекций (параметризованные запросы)
- CORS настройка для ограничения источников запросов
- Rate limiting для защиты от злоупотреблений
- Санитизация пользовательского ввода
- Безопасная обработка ошибок (не раскрываются внутренние детали)

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
| `LOG_LEVEL` | Уровень логирования (DEBUG, INFO, WARN, ERROR) | `INFO` |
| `CORS_ALLOWED_ORIGINS` | Разрешенные origins для CORS (через запятую) | `http://localhost:8000` |

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

# Проверка работы API
curl -X POST http://localhost:8080/api/feedback \
  -H "Content-Type: application/json" \
  -d '{"name":"Тест","phone":"+79991234567"}'

# Проверка данных в БД
docker-compose exec postgres psql -U postgres -d ecolife -c "SELECT * FROM applications ORDER BY created_at DESC LIMIT 5;"
```

## Лицензия

MIT
