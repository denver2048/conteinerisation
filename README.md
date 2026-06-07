## 🐍 Python app
Збираю в multi-stage build, щоб у майбутньому, якщо в додатку з’являться бібліотеки, які залежать від C/C++ залежностей, фінальний образ залишався максимально легким.

Також створюється користувач `pythonuser`, якому видані права `755` тільки для папки `/app`.

## Run locally
```bash
cd ~/conteinerisation
docker build -f ./app/Dockerfile -t python-app:slim ./app
docker run -d -p 8091:8090 python-app:slim
```

## Endpoints
- GET /
- GET /health

## ⚡ GO app
бираю в multi-stage build, тому що Go — це компільований мова програмування, і нам не потрібні вихідні файли проєкту в фінальному образі, а лише скомпільований бінарний файл.

Також створюється користувач з `UID` 101, оскільки в `scratch` образі неможливо створити іменованого користувача.

## Run locally
```bash
cd ~/conteinerisation
docker build -f ./go/Dockerfile -t go-app:scratch ./go
docker run -d -p 8090:8090 go-app:scratch
```

## Endpoints
- GET /hello

## 🔀 Проксування на різні порти
Обидва Docker-образи запускаються одночасно на одному хості завдяки використанню різних портів.

Це дозволяє паралельно запускати кілька сервісів без конфліктів портів.