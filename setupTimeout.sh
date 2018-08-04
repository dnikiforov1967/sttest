#!/bin/bash

curl -X PATCH -H 'Accept: application/json' -H 'Content-Type: application/json' http://localhost:8080/config/timeout/$1
