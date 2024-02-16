
build:
	go build -o bin/panmon  main.go

run:
	go run main.go


panmon:
	@./bin/panmon


panmon02:
	@./bin/panmon mon01 --dirToMonitor="./tmp" --dirTorecordJSON="./json" --concMan=4




# Testing of all libraries
test:
	@echo "================= testing a01"
	cd "./libs/a01" && go test -v

	@echo "================= test A02"
	cd "./libs/a02" && go test -v

	@echo "================= test A03"
	cd "./libs/a03" && go test -v

	@echo "================= testing Mon01"
	cd "./libs/mon01" && go test -v




#build2:
# GOARCH=amd64 GOOS=darwin go build2 -o bin/darwin/${BINARY_NAME}-darwin main.go
# GOARCH=amd64 GOOS=linux go build2 -o bin/linux/${BINARY_NAME}-linux main.go
# GOARCH=amd64 GOOS=windows go build2 -o bin/win/${BINARY_NAME}-windows main.go

#run: build2
# ./${BINARY_NAME}


# clean:
#  go clean
#  rm bin/darwin/${BINARY_NAME}-darwin
#  rm bin/linux/${BINARY_NAME}-linux
#  rm bin/win/${BINARY_NAME}-windows


