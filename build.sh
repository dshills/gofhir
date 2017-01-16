#!/bin/bash

ROOT=$GOPATH/src/github.com/dshills/gofhir
WATCHFILE=$1
FPATH=$(dirname "${WATCHFILE}")
FNAME=$(basename "${WATCHFILE}")
EXT=${FNAME##*.}

function build {
	BASEDIR=$1
	cd ${BASEDIR}
	go test && go vet && go install
}

function watch {
	# fswatch
	# -0 output nul seperator
	# -1 run just once
	# -r run recursive
	# xargs
	# -0 expect null sperator
	# -n 1 grab only the first arg
	# -I execute for each input line
	# {} replaced with args from fswatch

	# watch for file changes recursivly quit after first one is found and rerun the script
	fswatch -r ${ROOT} | xargs -n1 -I{} ${ROOT}/build.sh {}
}

if [ "${WATCHFILE}" == "" ]; then
	watch
else
	# Take action depending on the file type
	# Doing seperate builds if not in the main folder
	# allows code being worked on to be compiled even if
	# it is not used by the main application and as such would not be
	# compiled by simply doing go install in the root directory.
	# go install ./... works but test and vet get run on the vendor directory
	if [ "${EXT}" == "go" ]; then
		if [ "${FPATH}" == "${ROOT}" ]; then
			build ${FPATH}
		else
			build ${FPATH} && build ${ROOT}
		fi
	fi
fi
