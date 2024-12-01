# Тестовый проект с использованием mutex


Для фильтрации вывода (подсчета количества строк, содержащих "purchased"):
```bash
go run main.go | grep purchased | wc -l
```

Запуск программы 10 раз:
```bash
for run in {1..10}; do go run main.go | grep purchased | wc -l; done
```