#!/bin/sh

set -e
./main

./alist server || cat log/log.log
