#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

### Implementation ##################################################
function main {
  local -r version="1.0.000"
  local -r versionMajor=$(echo ${version} | cut -d. -f1)
  local -r versionMinor=$(echo ${version} | cut -d. -f2)

  printf "package finazon_grpc_go_client\n\nvar FINAZON_GRPC_HOST = \"grpc-v${versionMajor}-${versionMinor}.finazon.io:443\"" > ${SCRIPT_DIR}/pb/constants.go

  # generate service wrappers
  docker compose down --rmi local
  docker compose up
  docker compose down --rmi local
}

# Script's entry point: #############################################
main "${@}"
