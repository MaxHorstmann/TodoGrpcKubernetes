#!/bin/bash
docker run --rm -v $(pwd):/src $(docker build -q -f Dockerfile.local  .)
