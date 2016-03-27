#!/usr/bin/env bash

set -ex

curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "digraph{a->{1,2,3}}" > /tmp/g1.svg
curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "@examples/service.dot" > /tmp/g2.svg
