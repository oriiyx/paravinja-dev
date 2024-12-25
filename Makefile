#!/bin/zsh

run: build-app
	@./bin/main

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest
	@npm install -D @tailwindcss/typography

build-app:
	npx tailwindcss -i ./views/css/app.css -o ./public/styles.css
	@templ generate view
	@go build -o bin/main

