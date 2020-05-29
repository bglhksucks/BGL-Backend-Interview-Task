# BGL Take Home Task!!

Basically the application consists of a background task which fetches BTCUSD price from Binance.us every minute in MariaDB, and a RESTFUL API for queries of the followings from the database:
1. Last Price;
2. Price at a Time Point; and 
3. Average Price within a Time Range.

## Architecture
The application run on two docker containers:
1. Golang App (docker image built based on golang:1.14.2-alpine);
2. MariaDB (mariadb:10.5.1-bionic).

The two containers communicate through the docker network *bgl*.

## Code Organisation
Codes under the following folders focus on different functions:
###### *binance:*
Responsible for sending periodical requests to binance API to get the lastest BTCUSD prices, and storing the data into the database.

###### *db:*
Responsible for communication with the Database. In this case, it is MariaDB.

###### *helpers:*
Contains helper functions such as average value calculations, float64-string interconversions and timestamp formatting. 

###### *servers:*
Restful API codes that accept http requests and respond to them.
1. router.go: define routes of the API;
2. controllers.go: accept HTTP Requests, determine which functions to call and repond with appropriate HTTP Reponses;
3. responses.go: "bridges" between *controllers.go* and *db.go*.

## Build
1. Run `docker network create bgl` to docker network *bgl*.

2. Run the following command to start the DB container service:
```
docker run --net=bgl --name test-server_mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=bgl -e MYSQL_DATABASE=bgl -e MYSQL_USER=bgl -e MYSQL_PASSWORD=bgl -e MYSQL_DATABASE=bgl -d mariadb:10.5.1-bionic
```

3. Get the IP Address of DB container service with 
`docker inspect test-server_mariadb | grep IPAddress `

4. In .env, set `DB_CONNECT=<IP Address>:3306` with the IP Address obtained in the previous step.

5. In project root, run the following commands to start the application:
```
docker build -t test-server .
docker run -d --rm --name test-server --net=bgl -p 80:80 test-server
```

##### *(Alternative) Build with docker-compose*
1. In .env, set `DB_CONNECT=mariadb`.
2. In project root, run the following command to start the application:
```
docker-compose up -d
```
Make sure you have *docker* and *docker-compose* installed. Give about 30s - 1min for the system to initialize.

## Tests

In project root, run the following commands to run the tests:
```
go get "github.com/stretchr/testify/assert"
go get "github.com/DATA-DOG/go-sqlmock"
go test ./db/... ./helpers/... ./server/...
```

## CURL
### 1. Last Price
```
eg. curl http://localhost/lastprice
```
Example result:
```
{"lastPrice":"8853.23","time":"2020-05-12 15:54:47 +0800 HKT"}
```

### 2. Price at a Time Point
Endpoint: http://localhost/price/{timestamp}
```
eg. curl http://localhost/price/2020-05-13T15:23:00
```
Example result:
```
{"price":"8862.36","time":"2020-05-12 15:54:47 +0800 HKT"}
```

### 3. Average Price within a Time Range
Endpoint: http://localost/averagePrice/{from}/{to}
```
eg. curl http://localhost/averagePrice/2020-05-13T14:23:00/2020-05-13T15:23:00
```
Example result:
```
{"avgPrice":"8830.28","timeRange":"2020-05-13T03:23:00 to 2020-05-13T04:23:00"}
```