SHELL=/bin/bash
# build program
SRC_PATH = cmd
CFG_FILE = conf/config.yaml
BUILD_PATH = bin
SRC_FILES := $(shell cd $(SRC_PATH); find . -maxdepth 1 -type d|grep -v '/common')
TAGET := $(basename $(patsubst ./%,%,$(SRC_FILES)))
FILES := $(basename $(patsubst ./%,$(BUILD_PATH)/%,$(SRC_FILES)))

.PHONY: build $(TAGET) test-integration

# Example:
#   make build
#   make build GOFLAGS=-race
build:$(FILES)

$(TAGET):
	make -B $(BUILD_PATH)/$@

$(BUILD_PATH)/%: $(SRC_PATH)/%
	./scripts/build_script.sh $@ ./$^

.PHONY: clean
clean:
	rm -f $(FILES)

.PHONY: run
run:
	$(BUILD_PATH)/evolvestd -v -c $(CFG_FILE)


.PHONY: fmt
fmt:
	go mod tidy
	go fmt ./...
	golint ./...

.PHONY: test
test:
	go test -short -v ./...

.PHONY: gen-pb
gen-pb:
	protoc -I api/pb/evolvest/ api/pb/evolvest/evolvest.proto --go_out=plugins=grpc:api/pb/evolvest/


.PHONY: all
all: clean test build run


.PHONY: docker-build
docker-build:
	@docker build -t tenchael.com/evolvestd .

.PHONY: docker-run
docker-run:
	@docker run --name evolvestd -p 8762:8762 -p 8080:8080 tenchael.com/evolvestd


.PHONY: docker-clean
docker-clean:
	@docker container rm -f evolvestd
	@docker image rm -f evolvest_evolvestd_1 evolvest_evolvestd_2 evolvest_evolvestd_3
	@docker image rm -f tenchael.com/evolvestd

.PHONY: up
up:
	@docker-compose up

.PHONY: down
down:
	@docker-compose down