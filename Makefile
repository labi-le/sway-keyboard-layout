PROJ_NAME = sway-keyboard-layout

MAIN_PATH = main.go
BUILD_PATH = build/package/

build-default: clean
	go build --ldflags '-extldflags "-static" -s' -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

build-arm: clean
	GOOS=linux GOARCH=arm GOARM=7 make build-default

build-static: clean
	go build -ldflags "-w -linkmode external -extldflags "-static" -s" -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

build-debug: clean
	go build -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

clean:
	rm -rf $(BUILD_PATH)*

tests:
	go test -coverpkg=./... ./... -parallel=2