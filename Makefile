.PHONY: all
all: easyhide

.PHONY: easyhide
easyhide:
	go build -o $(PWD)/bin/easyhide $(PWD)/cmd/easyhide/...

.PHONY: fmt
fmt:
	go fmt $(PWD)/...

.PHONY: tidy
tidy:
	go mod tidy

bin/easyjson:
	GOBIN=$(PWD)/bin go install -ldflags '-s -w' github.com/mailru/easyjson/easyjson
