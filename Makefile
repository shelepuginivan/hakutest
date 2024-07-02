minify:
	./scripts/minify.sh

clean:
	go clean
	rm -r ${TARGET_DIR}
	rm hakutest hakuctl

test:
	go test -cover ./...
