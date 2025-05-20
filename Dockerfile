FROM caddy:builder AS builder
RUN xcaddy build \
    --with github.com/your-org/testcaddy/handler

FROM caddy:latest
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
COPY caddyfile /etc/caddy/Caddyfile
