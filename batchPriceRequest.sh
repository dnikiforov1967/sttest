#!/bin/bash

for (( i=1; i<=$1; ++i ))
do
	./priceRequest.sh
	sleep $2
done 
