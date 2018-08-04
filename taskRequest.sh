#!/bin/bash

curl -XGET -H 'Accept: application/json' --cookie "clientId=clientA" http://localhost:8080/price/$1
