dev:
	gowatch

clean:
	go clean
	rm -r ${TARGET_DIR}
	rm hakutest hakuctl

minify:
	./scripts/minify.sh

test:
	go test -cover ./...
