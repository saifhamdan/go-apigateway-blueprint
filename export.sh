# App env
export APP_ENV=development

# HTTP Config
export HTTP_HOST=127.0.0.1
export HTTP_PORT=8888
export HTTP_COOKIE_SECRET=your-most-secretive-secret


# Postgres Config
export POSTGRES_HOST=127.0.0.1
export POSTGRES_PORT=5432
export POSTGRES_USER=user
export POSTGRES_PASSWORD=123456
export POSTGRES_ROOT_PASSWORD=123456
export POSTGRES_NAME=apigateway

# Redis Config
export REDIS_HOST=127.0.0.1
export REDIS_PORT=6379

# NATS Config
export NATS_HOST=127.0.0.1
export NATS_PORT=4222

# Logs file path
export LOG_FILE=logs/api_gateway_blueprint.log

# Logs level
export LOG_LEVEL=debug

# Volumes path
export POSTGRES_VOLUME=/volumes/go-apigateway-blueprint/postgres
export REDIS_VOLUME=/volumes/go-apigateway-blueprint/redis