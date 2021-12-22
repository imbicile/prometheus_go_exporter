## Скелет для экспортера на Go
Планирую использовать для чтения своих метрик

Не используется
1. https://prometheus.io/docs/guides/go-application/
2. https://github.com/prometheus/client_golang

### Библиотеки
Справка и запуск с параметрами
```go
go get gopkg.in/alecthomas/kingpin.v2
```

### Отладка
```bash
go run go_exporter.go
```

### Сборка
```bash
go build go_exporter.go
```
### Запуск
```
./go_exporter
```
### Параметры
Name     | Description
---------|------------
-h | Справка
--listen-address="0.0.0.0" | Запуск на адресе
--listen-port="8080" | Запуск на порту
