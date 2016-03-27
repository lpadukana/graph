#!/usr/bin/env bash

set -ex

curl -v "http://localhost:9292/diagnostic/status/diagnosis"
