#!/bin/sh

for f in $*; do
  name=ruleguard-test
  echo "====== $f"
  docker build -t $name -f $f .
  docker run  --rm $name
  docker rmi $name > /dev/null
done

