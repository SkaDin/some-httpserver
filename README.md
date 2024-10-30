## Сокращатель ссылок (URL Shortener)

**Простой сокращатель ссылок, написанный на Go с использованием Viper для конфигурации и PostgreSQL для хранения данных.**

**Функциональность:**

* **Сокращает ссылки:** Преобразует длинные URL в уникальные короткие ссылки.
* **Кастомные короткие коды:** Позволяет задавать собственные короткие коды.
* **Перенаправление коротких ссылок:** Перенаправляет короткие ссылки на их исходные длинные URL.
* **Базовая аналитика:** Предоставляет базовую статистику для каждой короткой ссылки, например, количество кликов.
* **REST API:** Предоставляет API-интерфейсы для сокращения ссылок, получения исходных ссылок и просмотра аналитики.

**Технологии:**

* **Go:** Проект написан на Go.
* **Gorilla Mux:** Используется для маршрутизации HTTP-запросов.
* **Viper:** Используется для удобной загрузки конфигурации из различных источников (файлы, переменные окружения).
* **PostgreSQL:** Используется для хранения коротких ссылок и их соответствующих длинных URL.
* **GOOSE:** Используется для миграции базы данных PostgreSQL

**Начало работы:**

1. **Установите Go:** Убедитесь, что у вас установлен Go на вашей системе.
2. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/your-username/url-shortener.git
   ```

**Установите зависимости:**
```bash 
go mod tidy
```

**Запуск приложения:**
```bash
   docker-compose up build
```
**Настроить перменные окружения GOOSE:**
```bash export GOOSE_DRIVER=postgres
   export GOOSE_DBSTRING=postgresql://test:test@localhost:5432/testDB?sslmode=disable
```
**Запуск GOOSE:**
```bash
goose -dir db/migrations up
```

**Запуск гуся:**
```bash
goose -dir db/migrations up
```
