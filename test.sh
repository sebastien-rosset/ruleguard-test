#!/bin/sh

for f in Dockerfile*; do
  name=ruleguard-test
  echo "====== $f"
  docker build -q -t $name -f $f .
  docker run  --rm $name
  docker rmi $name > /dev/null
done

