BINARY_NAME=hakutest
BUILD_DIR=./build
SCRIPTS_DIR=./scripts
TARGET_DIR=./target

build: linux windows mac

linux:
	${SCRIPTS_DIR}/build_linux.sh

windows:
	${SCRIPTS_DIR}/build_windows.sh

mac:
	GOARCH=amd64 GOOS=darwin go build -trimpath -o ${TARGET_DIR}/mac/${BINARY_NAME} ./cmd/hakutest
	cp -r web ${TARGET_DIR}/mac
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-macos.tar.gz --transform 's/^./${BINARY_NAME}/' -C ${TARGET_DIR}/mac .
	rm -r ${TARGET_DIR}/mac

manual:
	mkdir -p ${TARGET_DIR}
	cp -r docs ${TARGET_DIR}/man
	gzip --recursive ${TARGET_DIR}/man
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-manual.tar.gz --transform 's/^./man/' -C ${TARGET_DIR}/man .
	rm -r ${TARGET_DIR}/man

clean:
	go clean
	rm -r ${TARGET_DIR}

test:
	go test -cover ./...
