## Prerequisite
`docker info`, демон запущен.
`minicube start`, локальный сервер запущен.

## Run
`./run.sh` в корне проекта соберет контейнеры и поднимет поды.
`kubectl get pods | grep ping`, `kubectl get pods | grep pong` - проверить что поды поднялись.
## Ping
`minicube service pong-service` вернет base url для доступа к сервису, эндпоинты:
`base-url/time`, `base-url/statistics`

