#!/usr/bin/env bash

TEXT_RED="\033[91m"
TEXT_GREEN="\033[92m"
TEXT_BLUE="\e[0;34m"
TEXT_WHITE="\033[0m"

## We want these values to 'pass', as in the program would consider them to be
## valid input.
tests_pass=('scaffold -account example-account'
            'scaffold -account example-account -region example-region'
            'scaffold -account example-account -region example-region -service example-service'
            'scaffold -account example-account -region example-region -service example-service -stack example-stack'
            'scaffold -account example-account -region example-region -service example-service -stack example-stack -kind example-kind'
            'scaffold -account example-account -region example-region -service example-service -stack example-stack -kind example-kind -resource example-resource'
            'scaffold -region example-region'
            'scaffold -region example-region -service example-service'
            'scaffold -region example-region -service example-service -stack example-stack'
            'scaffold -region example-region -service example-service -stack example-stack -kind example-kind'
            'scaffold -region example-region -service example-service -stack example-stack -kind example-kind -resource example-resource'
            'scaffold -service example-service'
            'scaffold -service example-service -stack example-stack' 
            'scaffold -service example-service -stack example-stack -kind example-kind' 
            'scaffold -service example-service -stack example-stack -kind example-kind -resource example-resource')
## We want these values to 'fail', as in the program would consider these to be 
## invalid input
tests_fail=('scaffold -account example-account -service example-service'
            'scaffold -account example-account -stack example-stack'
            'scaffold -account example-account -kind example-kind'
            'scaffold -account example-account -resource example-resource'
            'scaffold -region example-region -stack example-stack'
            'scaffold -region example-region -kind networking'
            'scaffold -region example-region -resource vpc'
            'scaffold -service example-service -kind example-kind'
            'scaffold -stack example-stack -resource example-resource'
            'scaffold -account example-account -resource example-resource' )

main () {
  test_scaffold
}

test_scaffold () {
  ## PASS is an exit code of 0
  printf "${TEXT_BLUE}VALID INPUT TEST${TEXT_WHITE}\n"
  for i in "${tests_pass[@]}"; do
    go run main.go ${i} -dry-run &>/dev/null
    if [[ ${?} != 0 ]]; then
      printf "\e[0;31m[FAIL] - Input: ${i}\e[0;37m\n"
    else
      printf "\e[0;32m[PASS] - Input: ${i}\e[0;37m\n"
    fi
  done
  ## PASS is an exit code of 1
  printf "${TEXT_BLUE}INVALID INPUT TEST${TEXT_WHITE}\n"
  for i in "${tests_fail[@]}"; do
    go run main.go ${i} -dry-run &>/dev/null
    if [[ ${?} != 1 ]]; then
      printf "\e[0;31m[FAIL] - Input: ${i}\e[0;37m\n"
    else
      printf "\e[0;32m[PASS] - Input: ${i}\e[0;37m\n"
    fi
  done

}

main
