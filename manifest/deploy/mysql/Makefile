DOCKERFILE = ./Dockerfile
# 禁止使用hub.docker.com 必须使用私有仓库
DOCKER_HUB = registry.cn-beijing.aliyuncs.com
NAMESPACE = hyhbackend
VERSION = 8.0
TAG = $(VERSION)
# 必须小写
IMG_NAME = adminmysql
IMG_FULL_NAME = $(DOCKER_HUB)/$(NAMESPACE)/$(IMG_NAME):$(TAG)


all: img

.PHONY : clean fmt build img help autoR test

img:
	docker build -t $(IMG_FULL_NAME) . -f $(DOCKERFILE)
	docker push $(IMG_FULL_NAME)

imgDebug:
	#docker run -it --rm --name admindb -e MYSQL_ROOT_PASSWORD="mysecretpassword" $(IMG_FULL_NAME)
	docker run -d --rm --name admindb -e MYSQL_ROOT_PASSWORD="mysecretpassword" $(IMG_FULL_NAME)
