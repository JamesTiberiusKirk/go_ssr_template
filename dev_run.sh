#!/bin/sh

reflex -d none -sr '.*\.(go)$' -- sh -c 'go run go_web_template'
