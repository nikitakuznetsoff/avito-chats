# avito-chats

## Запуск
Запуск осуществуляется с помощью `docker-compose up` из корня приложения
## Описание проекта

- Приложение запускается на порту `:9000`
- Использователя роутер из стандартной библиотекм
- В качестве хранилища используется `PostgreSQL`, которая содержит 4 таблицы:
    + `chats`, `messages`, `users` и `user_chat_relation` для отношения многие-ко-многим пользователей и чатов
    + Файл инициализации со структурой таблиц находится в папке `.sql`
- При неверном формате запроса возвращается ответ с соответствующим описанием и HTTP-ошибкой
- Для тестирования методов работы с БД использовал [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)

## Описание API с примерами запросов и ответов
**Добавление нового пользвателя**

Запрос
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username": "sample_name"}' \
  http://localhost:9000/users/add
```
Ответ
```
{"user_id": 1}
```
**Создать новый чат между пользователями**

- Идентификаторы должны быть int'ами

Запрос
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "chat_1", "users": [1,2]}' \
  http://localhost:9000/chats/add
```

Ответ
```
{"chat_id": 1}
```

**Отправить сообщение в чат от лица пользователя**

- Идентификаторы должны быть int'ами

Запрос
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1, "author": 1, "text": "hi"}' \
  http://localhost:9000/messages/add
```

Ответ
```
{"message_id": 1}
```

**Получить список чатов конкретного пользователя**

- Идентификатор должны быть int'ом

Запрос
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user": 1}' \
  http://localhost:9000/chats/get
```

Ответ: список чатов, отсортированный по времени создания последнего сообщения в чате
```
{"chats":[
    {"chat":1,"name":"chat_1","users":[1,2],"created_at":"2021-07-15T16:59:10.148576Z"},
    {"chat":2,"name":"chat_1","users":[1,3],"created_at":"2021-07-22T16:49:52.201781Z"}
  ]
}
```

**Получить список сообщений в конкретном чате**

- Идентификатор должны быть int'ом

Запрос
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"chat": 1}' \
  http://localhost:9000/messages/get
```

Ответ
```
{"messages":[
    {"id":3,"chat":1,"author":2,"text":"axcv","created_at":"2021-07-22T16:25:17.239599Z"},
    {"id":2,"chat":1,"author":2,"text":"asd","created_at":"2021-07-22T16:24:51.260359Z"},
    {"id":1,"chat":1,"author":1,"text":"hi","created_at":"2021-07-15T17:00:17.892784Z"}
  ]
}
```

## Структура проекта
Структура проекта в соответствии с [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
```
avito-chats
│   README.md
│   Dockerfile
│   docker-compose.yml
│
└───bin
│
└───.sql
│   |   db.sql
│
└───cmd
│   └───chatsapp
│       │   main.go
│   
└───internal
│   |
│   └───handlers
│   |   │   handler.go
│   |   │   chats.go
│   |   │   messages.go
│   |   │   users.go
│   |
│   └───models
│   |   │   chat.go
│   |   │   message.go
│   |   │   user.go
│   |
│   └───repository
│       │   repostiory.go
│       └───postgres
│       |   │   chat.go
│       |   │   message.go
│       |   │   user.go


```

- `bin/chatsapp` - бинарник для запуска проекта в контейнере
- `.sql/db.sql` - файл инициализации базы данных с созданием нужных таблиц
- `cmd/chatsapp/main.go` - файл для запуска приложения
- `internal/repository` - реализация работы с БД по паттерну "Репозиторий"
- `internal/handlers` - HTTP обработчики для запросов
- `internal/models` - описания сущностей БД
