# ios-backend
## Backend for IOS course in BMSTU Techpark
### Go, postgres, docker, Clean Architecture

Для запуска с заполеннием базы нужно создать файл .env  в корневой директории со следующими переменными:
```
COINMARKET_API_KEY=<ключ для CoinMarkert>
COINMARKET_URL=https://pro-api.coinmarketcap.com
FIATS=EUR,RUB,UAH,GBP,CNY 
```
FIATS - валюты в которые хотим 
конвертировать. USD - есть всегда, 
его указввать не нужно.