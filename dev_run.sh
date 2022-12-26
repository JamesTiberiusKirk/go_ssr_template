#!/bin/sh

export DEBUG=true
export DB_NAME=postgres
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASS=password
export HTTP_PORT=:3000
export HTTP_ROOT_PATH=/v1
export SESSION_SECRET=super_secret

reflex -d none -sr '.*\.(go|gohtml|env)$' -- go run go_ssr_template
