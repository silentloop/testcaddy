build:
	xcaddy build --with github.com/silentloop/testcaddy/handler=./handler --output ./caddy

run: build
	./caddy run --config ./caddyfile

clean:
	rm -f ./caddy
