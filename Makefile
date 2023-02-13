dev:
	export IMAGINE_USERNAME="test";export IMAGINE_PASSWORD="1234";go run ./main.go server

build:
	go build -o build/imagine main.go