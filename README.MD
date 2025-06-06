# 📖 LibraryAPI

Простое API для управления библиотечным каталогом на Go с использованием [Gin](https://github.com/gin-gonic/gin).

---

## Запуск проекта

```bash
go run main.go
```
Сервер поднимается на: http://localhost:8080


## Примеры запросов:
### GET /books
(Опциональный фильтр по автору):

```bash
GET /books?author=Толстой
```
Ответ:

```bash
[
  {
    "id": "3",
    "title": "Война и мир",
    "author": "Лев Толстой",
    "quantity": 6
  }
]
```
### Получить книгу по ID
```bash
GET /books/:id
```

```bash
GET /books/1
```
Ответ:

```bash
{
  "id": "1",
  "title": "Преступление и наказание",
  "author": "Фёдор Достоевский",
  "quantity": 3
}
```
### Добавить новую книгу
```bash
POST /books
```
Тело запроса:
```bash
{
  "id": "6",
  "title": "Новая книга",
  "author": "Новый Автор",
  "quantity": 10
}
```
Ответ:
```bash
{
  "id": "6",
  "title": "Новая книга",
  "author": "Новый Автор",
  "quantity": 10
}
```
### Выдать книгу (уменьшить quantity)
```bash
POST /books/checkout
```

Тело запроса:
```bash
{
  "id": "1"
}
```
Ответ:
```bash
{
  "id": "1",
  "title": "Преступление и наказание",
  "author": "Фёдор Достоевский",
  "quantity": 2
}
```
### Вернуть книгу (увеличить quantity)
```bash
POST /books/return
```
Тело запроса:
```bash
{
  "id": "1"
}
```
Ответ:

```bash
{
  "id": "1",
  "title": "Преступление и наказание",
  "author": "Фёдор Достоевский",
  "quantity": 3
}
```
# Разработчик: [Аринов Данияр](https://github.com/vegitobluefan)