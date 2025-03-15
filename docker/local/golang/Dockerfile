# Используем официальный образ Golang в качестве базового
FROM golang:1.24-alpine as builder

WORKDIR /app

ADD . /app

# Загружаем все зависимости
RUN go mod tidy

## Строим приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./app/cmd/app

# Используем более легкий образ для финального контейнера
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение из предыдущего контейнера
COPY --from=builder /app/main /app/main
COPY --from=builder /app/app/configs/local.yaml /app/local.yaml

# env
ENV CONFIG_PATH=local.yaml

# Открываем порт, на котором будет работать приложение
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
