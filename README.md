Application is served on localhost:8080.

To reproduce, just run `TEMPL_DEV_MODE=true templ generate -v --watch --cmd "go run ."`

To make testing different versions easier, I'm using a Makefile, just run the make command.

```bash
# Run using latest version of templ
TEMPL_VERSION=latest make test

# Run using last-working version
TEMPL_VERSION=v0.3.865 make test

# Run using version that introduced the bug
TEMPL_VERSION=v0.3.887 make test
```