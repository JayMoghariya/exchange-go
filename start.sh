docker build -t trading-system-go .

docker run -p 8080:8080 trading-system-go

docker-compose up --build

# docker-compose down