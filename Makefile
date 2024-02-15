
build:
	go build -o bin/fi  main.go

run:
	go run main.go


rfi:
	@./bin/fi

# modify separate arguments here.
rfi01:
	@./bin/fi f01



# modify separate arguments here.
rt01:
	go build -o bin/ti  main.go
	@./bin/ti t01

rt01-test:
	go test


test:

	@echo "================= testing a01"
	cd "./libs/a01" && go test -v

	@echo "================= test A02"
	cd "./libs/a02" && go test -v

	@echo "================= test A03"
	cd "./libs/a03" && go test -v

	@echo "================= testing T01"
	cd "./libs/t01" && go test -v

	@echo "================= test T02"
	cd "./libs/t02" && go test -v

	@echo "================= test F01"
	cd "./libs/f01" && go test -v

	@echo "================= test P01"
	cd "./libs/p01" && go test -v




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


