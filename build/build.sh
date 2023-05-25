#!/bin/bash
version=$1

if [ -z "$version" ]; then
   version="latest"
fi

basedir=`cd $(dirname $0); pwd -P`
echo ${basedir}
cd ${basedir}/../
sudo docker build -f ${basedir}/dockerfile -t registry.cnbita.com:5000/fusionhub/modelarts-adapter:$version .
