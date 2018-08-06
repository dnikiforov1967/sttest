# Test project

---

## Prerequisites

1. Database - SQLite 3
---
## How to build project

1. 

Download the project using go get tools:

```sh
go get github.com/dnikiforov1967/sttest
```

2.

Go to project root folder sttest and download the required go packages starting the following command:

```sh
./download-pkg.sh
```

You should download three Go packages:

* github.com/mattn/go-sqlite3
* github.com/gorilla/mux
* github.com/dnikiforov1967/accesslib
* github.com/stretchr/testify/assert

Package github.com/dnikiforov1967/accesslib implements rate limit functionality in the separate module

3. 

Compile the code:

```sh
./build.sh
``` 

4.

Project contains SQLite 3 database sttest.sqlt what already has required structures.
You can find structure descriptions in file create_struct.sql

If you want to re-create database structures, execute the script:

```sh
./createDb.sh
```

5.

Now you can start application using command:

```sh
./main
```
---
## How to configure application

The configuration file config.json is allocated in project root directory and has the following
structure

```sh
{
    "database":"sttest.sqlt",
    "timeout":6000,
    "limits":[
        {"clientId":"clientA","limit":7
        {"clientId":"clientC","limit":40}
    ]
}
```
* "database" value should NOT be changed
* "timeout" is the timeout of price calculation
* "limits" values show maximum number of requests to price API can issue each client per second

---

### Configuration API

To setup timeout you can use command

```sh
./setupTimeout.sh <timeout in milliseconds>
```
Request rate limitation can be setup for particular client using the command:

```sh
./setupRate.sh <client Id> <limit>
```

You can see current configuration using the request:

```sh
./getConfig.sh
```

---

## How to test application

### Go unit tests

To execute application unit tests you should call:

```sh
./goTest.sh
```

### Web service tests using CURL

#### Product service

##### 1. Create product

To create product you should execute POST request with JSON body containing product description
against /product URL. Please see insertData.sh script code for details.

You can modify the content of file 

```sh
body.json 
```

allocated in project folder in accordance to your wish and execute the command:

```sh
./insertData.sh
```

You should see the response similar to:

```json
{"name":"A","product_id":"ProductId1","category":"C1","quanto":true,"creationDate":"2018-08-03",
"expirationDate":"2018-01-02",
"terms":{"events":[
{"type":"CLOSE","terminal":true,"execution":
{"on":{"kind":"SCHEDULE"},"origin":"USA","type":"CLOSE"},
"cashDirection":{"path":"TO_INVESTOR","type":"CASH",
"payment":{"type":"X-TYPE","method":"X-METHOD","algorithmId":"X-ALGO"}}},
{"type":"EXECUTION","terminal":false,"execution":
{"on":{"kind":"IMMEDIATE"},"origin":"EUROPEAN","type":"EXECUTION"},
"cashDirection":{"path":"TO_BANK","type":"STOCK",
"payment":{"type":"Z-TYPE","method":"Z-METHOD","algorithmId":"Z-ALGO"}}}]}}
``` 

Take into account what product_id value uniquely identifies the product. You should use it 
in all other manipulations with data.

##### 2. Update product

To create product you should execute PUT request with JSON body containing product description
against /product/{id} URL where id is product_id of the product you want to change. 
Please see updateData.sh script code for details.

You can modify the content of file 

```sh
update.json 
```

allocated in project folder in accordance to your wish and execute the command:

```sh
./updateData.sh <original product id>
```

You should see the response like this

```sh
HTTP/1.1 200 OK
Date: Fri, 03 Aug 2018 15:46:50 GMT
Content-Length: 469
Content-Type: text/plain; charset=utf-8

{"name":"FX_Res_Knock_Into_FW_Imp_eu","product_id":"FX_Res_Knock_Into_FW_Imp_eu",
"category":"Tx-based OTC","quanto":true,"creationDate":"2018-05-25",
"expirationDate":"2020-05-25",
"terms":{"events":[{"type":"EXECUTION","terminal":true,
"execution":{"on":{"kind":"SCHEDULE"},"origin":"EUROPEAN","type":"EXECUTION"},
"cashDirection":{"path":"TO_INVESTOR","type":"CASH",
"payment":{"type":"TRIGGER_PRICE","method":"ALGORITHM","algorithmId":"FX_Res_Knock_Into_FW_Imp_eu"}}}]}}
```

##### 3. Fetch product

To select product execute the following script:

```sh
./selectData.sh <product id>
``` 

You should see the response like this

```sh
HTTP/1.1 200 OK
Date: Fri, 03 Aug 2018 15:46:50 GMT
Content-Length: 469
Content-Type: text/plain; charset=utf-8

{"name":"FX_Res_Knock_Into_FW_Imp_eu","product_id":"FX_Res_Knock_Into_FW_Imp_eu",
"category":"Tx-based OTC","quanto":true,"creationDate":"2018-05-25",
"expirationDate":"2020-05-25",
"terms":{"events":[{"type":"EXECUTION","terminal":true,
"execution":{"on":{"kind":"SCHEDULE"},"origin":"EUROPEAN","type":"EXECUTION"},
"cashDirection":{"path":"TO_INVESTOR","type":"CASH",
"payment":{"type":"TRIGGER_PRICE","method":"ALGORITHM","algorithmId":"FX_Res_Knock_Into_FW_Imp_eu"}}}]}}
```

##### 4. Delete product

To delete product execute the following script:

```sh
./deleteData.sh <product id>
``` 

You should see the response like this

```sh
HTTP/1.1 204 No Content
Date: Fri, 03 Aug 2018 16:20:06 GMT
```
---

#### Price Service

##### Asynchronious requests

###### * All requests to price API in scripts below executed under account clientA (using cookies).

To perform price calculation request you can use the floowing script:

```sh
./priceRequest.sh
```
This script posts to server the data from file price.json, allocated in root project directory.
Its content:

```sh
{
    "isin": "67462",
	"underlying": 177.82,
	"volatility": 0.818
}
```
You should immediately get the response like this

```sh
{"resource":"price/1"}
```
The last part of resource string is the unique number of price calculation task
You should use it to get the status of task using the request

```sh
./taskRequest.sh <task id>
```
where <task id> is the value returnd by price request
You get the response like
	
```sh
{"id":1,"isin":"67462","status":"IN PROGRESS","price":0,"date":""}
```
Status IN PROGRESS means what the price calculation is still being executed.
After 5 seconds you should get another response on the same request:

```sh
{"id":2,"isin":"67462","status":"COMPLETED","price":145456.76,"date":"2018-08-04T21:10:31+02:00"}
```
##### Waiting request

Alternatively you can use waiting API to get price. API shall wait for the completion of price calculation (
wait for the async procedure execution completion):

```sh
./priceRequestWait.sh

{"id":3,"isin":"67462","status":"COMPLETED","price":145456.76,"date":"2018-08-04T21:17:35+02:00"}
```

##### Timeouts

Setup timeout value equal to 2 seconds using command:

```sh
./setupTimeout.sh 2000
```
Start new price request:

```sh
./priceRequest.sh
{"resource":"price/2"}
```
After two seconds request the state of task. You shouls see TIMED OUT status. This means task execution was interrupted
by timeout

```sh
./taskRequest.sh 2
{"id":2,"isin":"67462","status":"TIMED OUT","price":0,"date":""}
```
Similarly start waiting request. You should get floowing result:

```sh
./priceRequestWait.sh
Task cancelled by timeout
```
Setup timeout value equal to 6 seconds:

```sh
./setupTimeout.sh 6000
```
Start new price request:

```sh
./priceRequest.sh
{"resource":"price/4"}
```
Before 6 seconds after start task request should return:

```sh
./taskRequest.sh 4
{"id":4,"isin":"67462","status":"IN PROGRESS","price":0,"date":""}
```
After 6 seconds you should get normal completion

```sh
./taskRequest.sh 4
{"id":4,"isin":"67462","status":"COMPLETED","price":145456.76,"date":"2018-08-05T20:24:03+02:00"}
```
Similarly:

```sh
./priceRequestWait.sh
{"id":5,"isin":"67462","status":"COMPLETED","price":145456.76,"date":"2018-08-05T20:26:59+02:00"}
```

##### Rate limit in action

To see the effect of the rate limit please use the following script

```sh
./priceRequestWait.sh 1000 0.01
```
The first parameter is the number of iterations in batch, the second one is the delay between sequential requests in seconds.
In example above we execute 100 iterations with 10 ms delay.

You should see the result like this:


```sh
{"resource":"price/4"}
{"resource":"price/5"}
{"resource":"price/6"}
{"resource":"price/7"}
{"resource":"price/8"}
{"resource":"price/9"}
{"resource":"price/10"}
Too many requests
Too many requests
Too many requests
Too many requests
Too many requests
Too many requests
...
Too many requests
Too many requests
{"resource":"price/11"}
{"resource":"price/12"}
{"resource":"price/13"}
{"resource":"price/14"}
{"resource":"price/15"}
{"resource":"price/16"}
{"resource":"price/17"}
Too many requests
Too many requests
Too many requests
...
```

Using config API you can manipulate rate limit value for clientA as well as delay value in command line and see different results

---




