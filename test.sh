#!/usr/bin/env bash

curl -G -v "http://localhost:9292/dot/to/svg" --data-urlencode "digraph test {a->b}"
