#!/bin/bash

for num in `seq 1 10`
do 
    curl -v -HHost:webservice.greeting.com http://192.168.64.2:30829/ | tail -1 >> ~/output.txt
done