#! /usr/bin/env bash

export CORE_PEER_ADDRESS="127.0.0.1:7052"
export CORE_CHAINCODE_ID_NAME="logisticscc:1.0"

go run .
