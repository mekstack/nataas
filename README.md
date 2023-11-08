# rpaas
Reverse proxy as a Service

# Envoy gRPC Клиент

Этот gRPC клиент предоставляет простой способ взаимодействия с Envoy, обратным прокси-сервером. Он предоставляет
возможность управления и настройки Envoy с помощью gRPC-интерфейса, а также распространяет важную информацию о состоянии
бэкенд-серверов.

## Флаги командной строки

Вам предоставляется несколько флагов командной строки для настройки клиента:

- `--debug`: Позволяет включить отладочное логирование сервера xDS.
- `--port`: Задает порт, на котором слушает xDS-сервер (по умолчанию: 18000).
- `--nodeID`: Определяет уникальный идентификатор узла (Node ID) для Envoy (по умолчанию: "test-id").
- `--cfgPath`: Указывает путь к файлу конфигурации (по умолчанию: "config/config.yaml").

## Функции и возможности

Этот gRPC клиент предназначен для взаимодействия с Envoy и предоставляет следующие функции:

1. **Перенаправление запросов (Request Forwarding):** Основная функция обратного прокси - обработка входящих HTTP
   запросов и направление их к соответствующим бэкенд-серверам.

2. **Балансировка нагрузки (Load Balancing):** Распределение входящего сетевого трафика между группой бэкенд-серверов
   для предотвращения перегрузки отдельных серверов.

3. **TLS Termination:** Расшифровка входящих TLS соединений на прокси и пересылка трафика на бэкенд-серверы внутри сети
   без шифрования для снижения вычислительной нагрузки на бэкенд-серверы.

## Поддерживаемые протоколы

Этот клиент поддерживает следующие версии HTTP:

- HTTP/2
- HTTP/1.1

## Запуск

Прежде чем запустить клиент, убедитесь, что у вас установлены все необходимые зависимости и выполните следующую команду:

```bash
go run main.go
envoy -c bootstrap-xds.yaml
