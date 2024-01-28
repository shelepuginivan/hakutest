BINARY_NAME=hakutest
TARGET_DIR=./target

build: linux windows mac

linux:
	GOARCH=amd64 GOOS=linux go build -o ${TARGET_DIR}/linux/${BINARY_NAME} ./cmd/hakutest
	cp -r web ${TARGET_DIR}/linux
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-linux.tar.gz --transform 's/^./${BINARY_NAME}/' -C ${TARGET_DIR}/linux .
	rm -r ${TARGET_DIR}/linux

windows:
	GOARCH=amd64 GOOS=windows go build -o ${TARGET_DIR}/${BINARY_NAME}/${BINARY_NAME}.exe ./cmd/hakutest
	cp -r web ${TARGET_DIR}/${BINARY_NAME}
	cd ${TARGET_DIR} && zip ${BINARY_NAME}-win64.zip -r ${BINARY_NAME}
	cp tools/* ${TARGET_DIR}/${BINARY_NAME}
	cd ${TARGET_DIR} && zip ${BINARY_NAME}-with-scripts-win64.zip -r ${BINARY_NAME}
	rm -r ${TARGET_DIR}/${BINARY_NAME}

mac:
	GOARCH=amd64 GOOS=darwin go build -o ${TARGET_DIR}/mac/${BINARY_NAME} ./cmd/hakutest
	cp -r web ${TARGET_DIR}/mac
	tar -czf ${TARGET_DIR}/${BINARY_NAME}-mac.tar.gz --transform 's/^./${BINARY_NAME}/' -C ${TARGET_DIR}/mac .
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

server:
	go build ./cmd/hakutest
	bash -c "trap 'rm ./hakutest' 2; ./hakutest server"

test:
	go test -c ./...
	for t in *.test; do \
		./$$t 1> /dev/null || { echo "TEST FAILED!"; rm *.test; exit 1; } ; \
	done
	rm *.test
