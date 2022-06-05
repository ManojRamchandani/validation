# awesome-api
awesome-api is a webserver that listens on the localhost @ port 9999
# Usage
It can be initialised, built, run, stoppped, cleaned using the following make targets
build   go.mod  Builds the webserver binary awesome-api
run      Starts the webserver awesome-api. logs stored in awesome-api.log
stop     Kills the webserver
clean    Deletes awesome-api binary & log file
test     Tests the webserver is up & running using curl calls
help     Prints the make targets available to invoke
