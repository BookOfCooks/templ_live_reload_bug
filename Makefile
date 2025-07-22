test:
	go install github.com/a-h/templ/cmd/templ@${TEMPL_VERSION}
	go mod edit -require=github.com/a-h/templ@${TEMPL_VERSION}
	go mod tidy
	TEMPL_DEV_MODE=true templ generate -v --watch --cmd "go run ."