credits:
	@echo "Made with ❤️ by [filirnd](http://github.com/filirnd) and [fedyfausto](http://github.com/fedyfausto)"
build:
	@make credits
	@go build -o ./bin/bigbrocore github.com/bigbroproject/bigbrocore/cmd
	@chmod +x ./bin/bigbrocore
	@echo "Build successfully!"
build-run:
	@make build
	@./bin/bigbrocore config/conf.yml