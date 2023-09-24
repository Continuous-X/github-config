#!/usr/bin/env bash
# default dependency-check path
DC_HOME=${DC_HOME:-'/c/DEV/tools/dependency-check/bin'}
PATH="${PATH}:${DC_HOME}/bin"
OUTPUT_PATH="target"

HTTP_PROXY_HOST="localhost"
HTTP_PROXY_PORT="3128"

dependency-check.sh --project github-config \
    -f HTML \
    --failOnCVSS 1 \
    --prettyPrint \
    --enableExperimental \
    -s . \
    -l ${OUTPUT_PATH}/depcheck.log \
    -o ${OUTPUT_PATH} \
    --proxyserver ${HTTP_PROXY_HOST} \
    --proxyport ${HTTP_PROXY_PORT} \
    --nonProxyHosts "" \
    --disableAssembly

#    --suppression dependency-check-suppression.xml \
