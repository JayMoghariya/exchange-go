- generate token using 
```bash
python3 -c "import secrets; print(secrets.token_hex(32))"
```

- create trader user using 
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"trader2","password":"Mystrong#00","role":"trader"}'
```

- sample env file
```env
# .env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=trading
DB_PORT=5432
JWT_SECRET=c9650526f927d0285b43832fbf9d2c7d552a80500493f4741018232ba7fc2931
JWT_EXPIRATION=24h
ADMIN_PASSWORD=Admin@123
ADMIN_USERNAME=admin
```