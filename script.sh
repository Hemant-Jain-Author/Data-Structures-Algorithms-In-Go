#!/bin/bash

#script to recursively travel a dir of n levels

function traverse() {
for file in "$1"/*
do
    if [ ! -d "${file}" ] ; then
	if [[ ${file} == *.go ]] ; then
		echo "${file} is go file"
		go build "${file}" &> "$1".log
		sleep 60 &
#	else
#        	# echo "${file} is a file"
	fi
    else
        echo "entering recursion with: ${file}"
        traverse "${file}"
    fi
done
}

function main() {
    traverse "$1"
}

main "$1"
