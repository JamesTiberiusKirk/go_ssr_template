#!/bin/sh

export DEBUG=true
export DB_NAME=go_web_template
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASS=password
export HTTP_PORT=:3000
export HTTP_ROOT_PATH=
export HTTP_ROOT_API_PATH=/api/v1
export SESSION_SECRET=super_secret

# reflex -d none -sr '.*\.(go|ts)$' -- sh -c 'npm run build && go run go_web_template'
reflex -d none -sr '.*\.(go)$' -- sh -c 'go run go_web_template'
