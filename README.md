## Геометрический решатель

Учебный проект по курсу "Разработка САПР"

## Запуск проекта

Запуск проекта происходит через *Docker* контейнеры посредством *docker-compose.yml*
```
docker-compose -f docker-compose.yml up -d --build
```
## Примеры работы с сервисом
(Тестовый примеры, чтобы проверить, живой ли сервер)

Отправка POST запроса :
```bash
curl --header "Content-Type: application/json" \
--request POST \
--data '{$FullJson$}' \
127.0.0.1:5001/api/v1/geomSolver/line/parallelism/
```
где вместо *$FullJson$* вставить полный json, выше показано как пример

Пример ожидаемого JSON в запросе:
```json
{
    "points": [
      {
        "X": 5.6,
        "Y": 6.9
      },
      {
        "X": 9.6,
        "Y": 6.9
      },
      {
        "X": 5.6,
        "Y": 13.9
      },
      {
        "X": 9.6,
        "Y": 18.9
      }
    ],
    "lines": [
      {
        "First": 0,
        "Second": 1
      },
      {
        "First": 2,
        "Second": 3
      }
    ],
    "equalTwoPoints": null,
    "distanceBetweenPoints": null,
    "FixationPoint": null,
    "BelongOfLine": null,
    "parallelTwoLines": null,
    "PerpenTwoLines": null,
    "CornerTwoLines": null,
    "VerticalLine": null,
    "HorizontLine": [
      0
    ]
}
```