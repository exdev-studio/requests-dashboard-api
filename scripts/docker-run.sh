docker run --rm --name requests-dashboard-api -d -it --env APP_LOG_LEVEL=debug --env=APP_BIND_ADDR=:8080 -p 8080:8080 exdev-studio/requests-dashboard-api:latest
