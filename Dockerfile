# Указываем базовый образ для сборки
FROM golang:1.23-alpine AS builder

# Устанавливаем необходимые зависимости для работы с Git
RUN apk add --no-cache git

# Создаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем остальные исходные файлы
COPY . .

# Собираем приложение
RUN go build -o calendar-app ./cmd/main.go

# Используем минимальный образ для запуска собранного приложения
FROM alpine:3.16

# Создаем рабочую директорию для приложения
WORKDIR /app

# Копируем скомпилированный бинарник из предыдущего шага
COPY --from=builder /app/calendar-app .

# Устанавливаем часовой пояс, если необходимо
RUN apk add --no-cache tzdata
# Указываем порт, который будет использовать приложение
EXPOSE 8080
# Указываем команду для запуска приложения
CMD ["./calendar-app"]


