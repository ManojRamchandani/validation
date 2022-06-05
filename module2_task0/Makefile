.PHONY_TARGET: build init run stop clean test help
.DEFAULT_GOAL := help

build: go.mod ## Builds the webserver binary awesome-api
	echo "Building awesome-api bin" 
	go build 

go.mod: # Initialses the webserver
	echo 'Initialisin moduled awesome-api'
	go mod init github.com/ManojRamchandani/validation/module2_task0/awesome-api && \
	go get "github.com/gorilla/mux" && \
	go mod tidy 
	

run: ## Starts the webserver awesome-api. logs stored in awesome.log
	./awesome-api >./awesome.log 2>&1 &

stop: ## Kills the webserver
	kill "$(shell pgrep awesome-api)"

clean: ## Deletes awesome-api binary & log file
	rm -f ./awesome-api
	rm -f ./awesome.log

test: ## Tests the webserver is up & running using curl calls
	curl http://localhost:9999 
	curl http://localhost:9999/health

help: ## Prints the make targets available to invoke
	echo '# awesome-api' > ./Readme.md
	echo 'awesome-api is a webserver that listens on the localhost @ port 9999' >> ./Readme.md
	echo '# Usage' >> ./Readme.md
	echo 'It can be initialised, built, run, stoppped, cleaned using the following make targets' >> ./Readme.md
	sed -n 's/##\([^\(]\)/ /p' ./Makefile 	\
	| column -t -s ':'
	sed -n 's/##\([^\(]\)/ /p' ./Makefile 	\
	| column -t -s ':' >> ./Readme.md
