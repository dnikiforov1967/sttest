# Test project

## Prerequisites

1. Database - SQLite 3.X

## How to build project

1. 

Download the project using git or go get tools:

```
git clone https://github.com/dnikiforov1967/sttest.git
```

2.

Go to project root folder sttest and download the required go packages starting the following command:

```
./download-pkg.sh
```

You should download three Go packages:

github.com/mattn/go-sqlite3
github.com/gorilla/mux
github.com/dnikiforov1967/accesslib

Package github.com/dnikiforov1967/accesslib implements rate limit functionality in the separate module

3. 

Compile the code:

```
./build.sh
``` 

4.

Project contains SQLite 3 database sttest.sqlt what already has required structures.
You can find structure descriptions in file create_struct.sql

If you want to re-create database structures, execute the script:

```
./createDb.sh
```

5.

Now you can start application using command:

```
./main
```


## How to test application

### Go unit tests

To execute application unit tests you should call:

```
./goTest.sh
```

### Web service tests using CURL
