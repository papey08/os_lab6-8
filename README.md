# os_lab6-8

Этот проект является изменённой версией [лабораторной работы №6-8 по курсу 
«Операционные системы»](https://github.com/papey08/MAI_skat/tree/main/Operation_Systems/os_lab6-8).

### Отличия

|                       | os_lab6-8v2021         | os_lab6-8v2023          |
|-----------------------|------------------------|-------------------------|
| Язык                  | C++                    | Golang                  |
| Брокер сообщений      | ZeroMQ                 | RabbitMQ                |
| Способ хранения узлов | Бинарное дерево поиска | Красно-чёрное дерево    |

## Запуск и работа

***Для запуска должны быть установлены go 1.19 и Docker***
### Запускаем `rabbitmq` с помощью образа Docker
```
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```
### Запускаем `consumer.go`
```
go run cmd/consumer.go
```
### В другом терминале запускаем `publisher.go`
```
go run cmd/publisher.go
```
### `publisher.go` поддерживает следующие форматы сообщений
* `insert 1` — создать таймер с ключом 1
* `start 1` — запустить таймер 1
* `get 1` — узнать время на таймере 1

`consumer.go` выводит результаты сообщений

## Описание проекта

* **consumer.go** — в структуре map из стандартной библиотеки golang (в планах 
заменить на красно-чёрное дерево) по целочисленному ключу хранятся узлы 
сообщений. Каждый узел содержит объект `TTimer`, у таймера есть два метода: 
`Start` (запускает таймер) и `GetTime` (возвращает количество секунд с момента 
запуска). К каждому таймеру можно обратиться по ключу и применить любой метод. 
Команды программа получает от `publisher` через очередь сообщений.
* **publisher.go** — принимает из стандартного ввода команды и передаёт их в 
очередь сообщений. Возможен запуск нескольких экземпляров `publisher`, при 
этом `consumer` будет обрабатывать все полученные сообщения.

## Roadmap

* [x] Создать интерфейс `DataStructure` со всеми основными операциями для структур данных (вставка, поиск, удаление)
* [x] Реализовать структуру `TTimer` и методы `Start` и `GetTime`
* [x] Покрыть `TTimer` тестами
* [x] Добавить очередь сообщений
* [x] Реализовать красно-чёрное дерево (поиск и вставку)
* [x] Покрыть красно-чёрное дерево тестами
* [x] Реализовать метод `Pause` для `TTimer`
* [ ] Реализовать удаление из красно-чёрного дерева
* [ ] Добавить файл с конфигами
