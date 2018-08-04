#!/bin/bash

curl -XPOST -H 'Accept: application/json' -H 'Content-Type: application/json' --data @body.json http://localhost:8080/product

