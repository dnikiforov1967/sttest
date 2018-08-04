# Test project

---

## Prerequisites

1. Database - SQLite 3.X
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
"database" value should NOT be changed
"timeout" is the timeout of price calculation
"limits" values show maximum number of requests to price API can issue each client per second

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

To perform price calculation request you can use the floowing script:

```sh
./priceRequest.sh
```


---




