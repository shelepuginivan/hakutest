BINARY_NAME=hakutest
BUILD_DIR=./build
SCRIPTS_DIR=./scripts
TARGET_DIR=./target

build: linux windows mac

linux:
	${SCRIPTS_DIR}/build_linux.py

windows:
	${SCRIPTS_DIR}/build_windows.sh

mac:
	${SCRIPTS_DIR}/build_macos.py

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
