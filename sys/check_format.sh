#!/bin/bash

TMP_DIR=$(pwd | rev | cut -d '/' -f-1 | rev)

if [ $(whoami) = "root" ]; then
  echo ""
  echo "No es recomendable que se corra como root"
  echo "Presione Ctrl-C para salir o espere 10s"
  sleep 10
fi

if [ $TMP_DIR = "sys" ]; then
  cd ..
else if [ $TMP_DIR != "tools" ]; then
    echo "Error"
    exit 1
  fi
fi

cd src
go fmt
if [ $(which golint) = "golint not found" ]; then
  echo "Instala golint :D"
fi
