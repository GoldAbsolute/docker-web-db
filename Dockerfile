# syntax=docker/dockerfile:1

# Из образа для go
FROM golang:latest
# Имя рабочей директории, мб
WORKDIR /app
# Сначало копируются файлы с зависимостями
COPY go.mod ./
COPY go.sum ./
# Исполнение команд для скачивания зависимостей
RUN go mod download
RUN go mod tidy
# Копирование всех остальных файлов
COPY . ./
# Подключение зависимостей которые ущербные и не вошли в общий сборник
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
# docker-go-app - название бинарника
RUN go build -o /docker-go-app
# Контернер, выставляемые порты
EXPOSE 80
# Указывает входную точку, мб, в данном случае путь до собираемого бинарника
CMD [ "/docker-go-app" ]