#!/usr/bin/env bash

gox -osarch="linux/amd64" --output="build/app"


#docker run -p 5005:80 --env PORT=80 test