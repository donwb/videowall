#
# The include should be a single file that contains:
# export APIKEY := {APIKEY}
# export SECRET := {SECRET}
#
include env

$(info $$APIKEY is [${APIKEY}])
$(info $$SECRET is [${SECRET}])

all:
	go run *.go

build: ## Build
	go build *.go

test: 
	gow -e=go,mod,html,js run .
	
run:
	gin

docker-build:
	docker build -t donwb/videowall:0.7 .


docker-run:
	docker run --env-file=docker-env -p 80:1323 -it donwb/videowall:0.7 

.DEFAULT_GOAL := all