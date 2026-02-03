FROM python:3.11-alpine

WORKDIR /app

# Копируем статические файлы
COPY . .

EXPOSE 8000

# Запускаем простой HTTP сервер
CMD ["python3", "-m", "http.server", "8000"]

