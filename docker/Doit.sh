#!/bin/bash
docker build -t "azwebapp:Dockerfile" . 
docker tag azwebapp:Dockerfile fz96tw/azwebapp:Dockerfile
docker push fz96tw/azwebapp:Dockerfile


