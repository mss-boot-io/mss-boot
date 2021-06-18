#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

#-------------------------------------------------------------------------------
# gen proto

function mss::proto::gen() {
    local projects=( $1 common error )
    local type=$2
    local rootpath=$3
    local proto_json=$4
# local parent_rootpath=$(dirname "${rootpath}")

  out_args=""
  if [ "${type}" == "go" ]; then
    out_args=" -I ${rootpath}/src/go/vendor \
    --go_out=plugins=grcp,paths=source_relative:${rootpath}/src/go
    "
  else
    echo "not supported"
    # shellcheck disable=SC2242
    exit -1
  fi

# mkdir -p ${rootpath}/src/go/.out
  for service in "${projects[@]}"; do
    echo "begin to compile service:[$service]"
    # shellcheck disable=SC2045
    for sub in `ls ${service}`; do
      local walk_dir="$service/$sub"
      if [ ! -d $walk_dir ]; then
        walk_dir="$service"
      fi
      for pf in $walk_dir/*.proto; do
        if [ ! -f "$pf" ]; then continue; fi
          protoc -I $GOPATH/src/ \
            -I ${rootpath} \
            ${out_args} ${proto_json} \
            ${rootpath}/api/protobuf/$pf
        if [ $? -ne 0 ]; then
          echo "compile: [$pf] failed"; exit -1;
        else
          echo "compile: [$pf] success";
        fi
      done
    done
  done
}
