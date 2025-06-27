#!/usr/bin/env bash

function run() {
    local script="./scripts/$1.sh"
    if [[ -x "$script" ]]; then
        "$script"
    else
        echo "Command '$1' not found"
    fi
}

while :; do
    printf "stratforge> "
    read input
    [[ "$input" == "exit" ]] && break
    run "$input"
done
