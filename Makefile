APP = admin
BINARY = main
GF = gf
OUT_PATH = ./temp
DOCKERFILE = ./manifest/docker/Dockerfile
# 禁止使用hub.docker.com 必须使用私有仓库
DOCKER_HUB = registry.cn-beijing.aliyuncs.com
NAMESPACE = hyhbackend
VERSION = 0.0.1.1
TAG = $(VERSION)
# 必须小写
IMG_NAME = admin
IMG_FULL_NAME = $(DOCKER_HUB)/$(NAMESPACE)/$(IMG_NAME):$(TAG)


all: build

.PHONY : clean fmt build img help autoR test

clean:
	#go clean -i $(GO_FLAGS) $(SOURCE_DIR)
	rm -f ./$(BINARY)
	rm -f ./$(OUT_PATH)/linux_amd64/$(BINARY)
	rm -rf ./$(OUT_PATH)/$(VERSION)
	docker images | grep $(IMG_NAME) | sort | awk '{print $3}' | xargs docker rmi

cleanAll:
	#go clean -i $(GO_FLAGS) $(SOURCE_DIR)
	rm -f ./$(BINARY)
	rm -rf ./OUT_PATH/*

fmt:
	goimports -w ...

build:
	$(GF) build
	cp $(OUT_PATH)/$(VERSION)/linux_amd64/main $(OUT_PATH)/linux_amd64/main
	file $(OUT_PATH)/linux_amd64/main

img:
	make build
	docker build -t $(IMG_FULL_NAME) . -f $(DOCKERFILE)
	docker push $(IMG_FULL_NAME)

imgR:
	docker run  --rm -p 8199:8199 \
	-v "`pwd`/manifest/config/config_docker.yaml":/app/config/config.yaml \
	-v "`pwd`/manifest/config/rbac_model.conf":/app/config/rbac_model.conf \
	-v "`pwd`/dist/":/app/dist/ \
	$(IMG_FULL_NAME)

imgDebug:
	docker run -it --rm -p 8199:8199 \
	-v "`pwd`/manifest/config/config_docker.yaml":/app/config/config.yaml \
	-v "`pwd`/manifest/config/rbac_model.conf":/app/config/rbac_model.conf \
	-v "`pwd`/dist/":/app/dist/ \
	$(IMG_FULL_NAME) /bin/bash

help:
	@echo "make docker - 编译镜像!"
	@echo "make autoR - tx环境,自动部署!"

test:
	go test -v ./tsts
	#go test -v ./testsdb
	go test -v ./testsdb/lib_file_test.go
	go test -v ./testsdb/redis_test.go
	go test -v ./testsdb/service_gps_test.go
