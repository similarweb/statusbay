#!/bin/bash
set -e

# Validate it

if [ -z "${MODE}" ]; then
    echo "Please set the \$MODE environment variable" 1>&2
    exit 1
fi

/bin/statusbay -mode=${MODE}