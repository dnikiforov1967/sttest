#!/bin/bash

curl -XDELETE -H'Accept: application/json' -D- http://localhost:8080/product/$1
