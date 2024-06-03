BINARY_NAME=hakutest
BUILD_DIR=./build
SCRIPTS_DIR=./scripts
TARGET_DIR=./target

build: linux windows mac

linux:
	${SCRIPTS_DIR}/build_linux.sh

windows:
	GOARCH=amd64 GOOS=windows go build -trimpath -o ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}.exe ./cmd/hakutest
	GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags "-H=windowsgui" -trimpath -o ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}-server.exe ./cmd/hakutest-server
	GOARCH=amd64 GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags "-H=windowsgui" -trimpath -o ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}-statistics.exe ./cmd/hakutest-statistics
	cp -r web ${TARGET_DIR}/${BINARY_NAME}
	cd ${TARGET_DIR}/${BINARY_NAME} && ../../scripts/generate-wxs.py
	cp assets/${BINARY_NAME}.ico ${TARGET_DIR}/${BINARY_NAME}
	cd ${TARGET_DIR}/${BINARY_NAME} && wixl ${BINARY_NAME}.wxs -o ../${BINARY_NAME}-win64.msi
	rm ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}.wxs ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}.ico
	cd ${TARGET_DIR} && zip ${BINARY_NAME}-win64.zip -r ${BINARY_NAME}
	rm -r ${TARGET_DIR}/${BINARY_NAME}

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
