#
# The include should be a single file that contains:
# export APIKEY := {APIKEY}
# export SECRET := {SECRET}
#
include env

$(info $$APIKEY is [${APIKEY}])
$(info $$SECRET is [${SECRET}])
$(info $$CONN is [${CONN}])

all:
	go run *.go

build: ## Build
	go build *.go

test: 
	gow -e=go,mod,html,js,css run .
	
run:
	gin

docker-buildM1:
	docker build --platform linux/amd64 -t donwb/videowall:0.8 .


docker-build:
	docker build -t donwb/videowall:0.8 .

docker-run:
	docker run --env-file=docker-env -p 80:1323 -it donwb/videowall:0.8

.DEFAULT_GOAL := all