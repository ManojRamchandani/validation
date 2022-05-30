#!/bin/bash
# Installing gohugo & make
apt-get update && apt-get install -y hugo make
make build
