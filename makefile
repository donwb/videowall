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
	gow -e=go,mod,html run .
	
run:
	gin

docker-build:
	docker build -t videowall .


docker-run:
	docker run -p 3000:3001 -it videowall 

.DEFAULT_GOAL := all