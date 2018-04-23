#!/bin/bash

TEXT=$@

if [ -z $1 ]; then
  echo "enter a fucking quote in quotes"
  exit
fi

QUOTE="{\"text\": \"$TEXT\"}"
echo $QUOTE
curl -H "Content-Type: application/json" -d "$QUOTE" http://localhost:8080/this/is/for/alex/only/to/create
