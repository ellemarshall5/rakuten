BINARY_NAME=rakuten.out
	

build:
	go build -o ${BINARY_NAME} rakuten.go


run:
	go build -o ${BINARY_NAME} rakuten.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
