BINARY_NAME=myazwebapp
build:
	env GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}_linux_amd64 main.go 
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	cp ${BINARY_NAME}_linux_amd64 docker/.

run: build

clean:
	rm ${BINARY_NAME}_linux_amd64
	rm ${BINARY_NAME}_darwin_amd64
