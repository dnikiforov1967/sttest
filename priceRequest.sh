#!/bin/bash

curl -XPOST -H 'Accept: application/json' -H 'Content-Type: application/json' --cookie "clientId=clientA" --data @price.json http://localhost:8080/price
