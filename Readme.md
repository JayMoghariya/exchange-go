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

