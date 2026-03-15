# secure-app — Demo CI Project

Демонстрационный проект для показа работы CI-инструментов.

## Намеренные ошибки в коде

| Файл | Ошибка | Инструмент |
|------|--------|-----------|
| main.go:11 | Hardcoded password | gosec G101 |
| main.go:15 | MD5 для хэширования | gosec G401 |
| main.go:21 | Игнорирование ошибки os.Create | errcheck |
| main.go:22 | Игнорирование ошибки WriteString | errcheck |
| math_test.go:16 | Тест с неправильным ожидаемым значением | go test |
| go.mod | Устаревшая версия golang.org/x/crypto | govulncheck |

## Как воспроизвести

### 1. Запустить линтер
```
ci lint
```
Ожидаемый результат: ошибки gosec и errcheck

### 2. Запустить тесты
```
ci test
```
Ожидаемый результат: FAIL в TestDivide

### 3. Проверить уязвимости
```
ci vuln
```
Ожидаемый результат: уязвимость в golang.org/x/crypto

### 4. Попробовать закоммитить (сработает pre-commit)
```
git add .
git commit -m "add feature"
```
Ожидаемый результат: pre-commit заблокирует коммит с ошибками линтера

## Исправление

Чтобы исправить и сделать коммит успешным:
1. Убрать hardcoded password
2. Использовать bcrypt вместо MD5
3. Обрабатывать все ошибки
4. Исправить тест TestDivide (6 → 5)
5. Обновить golang.org/x/crypto до актуальной версии
