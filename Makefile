BINARY_NAME=hakutest
TARGET_DIR=./target

build: linux windows mac

linux:
	GOARCH=amd64 GOOS=linux go build -o ${TARGET_DIR}/linux/${BINARY_NAME} .
	cp -r web ${TARGET_DIR}/linux
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-linux.tar.gz --transform 's/^./${BINARY_NAME}/' -C ${TARGET_DIR}/linux .
	rm -r ${TARGET_DIR}/linux

windows:
	GOARCH=amd64 GOOS=windows go build -o ${TARGET_DIR}/win64/${BINARY_NAME}.exe .
	cp -r web ${TARGET_DIR}/win64
	cd ${TARGET_DIR} && mv win64 ${BINARY_NAME} && zip ${BINARY_NAME}-win64.zip -r ${BINARY_NAME}
	rm -r ${TARGET_DIR}/${BINARY_NAME}

mac:
	GOARCH=amd64 GOOS=darwin go build -o ${TARGET_DIR}/mac/${BINARY_NAME} .
	cp -r web ${TARGET_DIR}/mac
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-mac.tar.gz --transform 's/^./${BINARY_NAME}/' -C ${TARGET_DIR}/mac .
	rm -r ${TARGET_DIR}/mac

clean:
	go clean
	rm -r ${TARGET_DIR}

test:
	go test ./...
