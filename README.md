<img width="515" height="40" alt="image" src="https://github.com/user-attachments/assets/f4a845cc-2c47-4f03-9953-6cf985e3d507" /># Практическое задание №11

## Тема

GraphQL API для домена `Task` на базе `gqlgen`.

## Цель работы

Реализовать GraphQL API для задач и показать отличие GraphQL-подхода от REST на знакомом домене `Task`.

## Основа работы

Проект использует тот же домен `Task`, что и REST/Redis-сервис из `pz9-redis-cache`. Чтобы проект запускался самостоятельно, хранилище реализовано внутри `internal/store`.

## Требования к выполнению

- описать GraphQL-схему;
- сгенерировать код через `gqlgen`;
- реализовать query и mutation для задач;
- добавить in-memory-хранилище;
- запустить GraphQL endpoint и playground.

## Что реализовано

- В `graph/schema.graphqls` описаны тип `Task`, `Query` и `Mutation`.
- Сгенерированы файлы gqlgen.
- В `internal/store` реализовано in-memory-хранилище.
- В `graph/schema.resolvers.go` реализованы резолверы чтения, создания, обновления и удаления задач.
- В `cmd/graphql/main.go` настроены GraphQL playground и endpoint `/query`.

## Структура проекта

```
pz11-graphql/
├── cmd/
│   └── graphql/
│       └── main.go
├── graph/
│   ├── model/
│   ├── generated.go
│   ├── resolver.go
│   ├── schema.graphqls
│   └── schema.resolvers.go
├── internal/
│   └── store/
├── BASED_ON_PREVIOUS.md
├── gqlgen.yml
├── go.mod
├── go.sum
└── README.md
```

## Используемые технологии

- Go.
- GraphQL.
- `github.com/99designs/gqlgen`.
- GraphQL schema.
- Resolver.
- In-memory store.

## Как работает

Клиент отправляет GraphQL-запросы на единый endpoint `/query`. В отличие от REST, клиент сам выбирает нужные поля ответа. Резолверы принимают запрос, обращаются к in-memory-хранилищу и возвращают данные модели `Task`.

## Генерация gqlgen

Если менялась схема, нужно повторно сгенерировать код:

```
go run github.com/99designs/gqlgen generate
```

## Запуск

```powershell
go run .\cmd\graphql
```
<img width="515" height="40" alt="image" src="https://github.com/user-attachments/assets/75dbf4fc-7319-4ddc-beab-bc14c1954288" />


После запуска открыть: `http://localhost:8080/`

<img width="915" height="779" alt="image" src="https://github.com/user-attachments/assets/c09d20a5-7d6c-4374-a81d-dbd719d01ab2" />



## Проверка

```graphql
query {
  tasks {
    id
    title
    done
  }
}
```

<img width="1832" height="531" alt="image" src="https://github.com/user-attachments/assets/da6974ed-782c-4c73-af10-3c8bd32cdbbf" />


## Ожидаемый результат

GraphQL playground открывается в браузере, запрос `tasks` возвращает список задач, а mutation позволяют создавать, изменять и удалять задачи.

## Вывод

В ходе работы был реализован GraphQL API на Go. Практика показывает, что GraphQL дает клиенту возможность точнее управлять структурой ответа и использовать один endpoint для разных операций.
