#! /usr/bin/env bash

RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m'

image_tag="gohackathon"
port=5432
db_name="hackathon_dev"
docker_file="dockerfile"

printOk() {
   echo -e "$BLUE $1 $NC"
}

printBad() {
   echo -e "$RED $1 $NC"
}

checkDB() {
   FILE=database.go
   if test -f "$FILE"; then
      printOk "$FILE exists"
   else
      printBad "$FILE not found"
   fi

   go run $FILE
   if [ $? -eq 0 ]; then
      printOk "DB is Successfully setup, Happy Coding"
   else
      printBad "Something went wrong, try running $FILE file manually"
   fi
}

createDockerFile() {
   touch $docker_file
   echo "FROM postgres:latest" >$docker_file
   if [ $? -eq 0 ]; then
      printOk "Docker File created successfully"
   else
      printBad "Couldnt create docker file"
      exit 1
   fi
}

createImage() {
   docker build -f dockerfile -t $image_tag .
   if [ $? -eq 0 ]; then
      printOk "Image Built Successfully"
   else
      printBad "Something went wrong, try again"
   fi
}

docker rmi $image_tag
if [ $? -eq 0 ]; then
   printOk "Image found, removing image from system"
else
   printBad "Image not found, creating new image"
fi

if test -f "$docker_file"; then
   createImage
else
   printBad "Docker file not found, creating docker file $docker_file"
   createDockerFile
   createImage
fi

docker run --rm -d -p $port:$port -e POSTGRES_DB=$db_name $image_tag:latest
if [ $? -eq 0 ]; then
   printOk "Container is running, checking for successful DB setup, Plase wait......."
   sleep 6
   checkDB
else
   printBad "Couldnt start up the container, please try again"
fi
