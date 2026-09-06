# cloud-file-storage-go
Многопользовательское файловое облако. Пользователи сервиса могут использовать его для загрузки и хранения файлов

## Миграции
Использована библиотека - golang-migrate

### Создать миграцию

```bash
make migration name=create_users
```

Создаст:

```text
000001_create_users.up.sql
000001_create_users.down.sql
```

### Применить миграции

```bash
make migrate-up
```

### Откатить последнюю

```bash
make migrate-down
```

### Посмотреть версию

```bash
make migrate-version
```

### Исправить dirty state

```bash
make migrate-force version=1
```
