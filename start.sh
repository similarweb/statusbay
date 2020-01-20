#!/bin/bash
set -e

# Validate it
if [ -z "${PREFIX}" ]; then
    echo "Please set the \$PREFIX environment variable" 1>&2
    exit 1
fi

if [ -z "${CONSUL_ADDR}" ]; then
    echo "Please set the \$CONSUL_ADDR environment variable" 1>&2
    exit 1
fi

if [ -z "${MODE}" ]; then
    echo "Please set the \$MODE environment variable" 1>&2
    exit 1
fi

if [ -z "${VAULT_ROLE_ID}" ]; then
    echo "Please set the \$VAULT_ROLE_ID environment variable" 1>&2
    exit 1
fi

if [ -z "${VAULT_ADDR}" ]; then
    echo "Please set the \$VAULT_ADDR environment variable" 1>&2
    exit 1
fi

vault_token=$(curl --request POST --data '{"role_id": "'${VAULT_ROLE_ID}'"}' ${VAULT_ADDR}/v1/auth/approle/login | jq .auth.client_token | tr -d '"')
if [ $vault_token != "" ]; then        
    echo INFO "Exporting VAULT_TOKEN: ${vault_token} & VAULT_ADDR: ${VAULT_ADDR}"
    export VAULT_TOKEN="$vault_token"
fi


if [ "${MODE}" == "kubernetes" ]; then

    if [ -z "${REGION}" ]; then
        echo "Please set the \$REGION environment variable" 1>&2
        exit 1
    fi

    if [ -z "${ENVIRONMENT}" ]; then
        echo "Please set the \$ENVIRONMENT environment variable" 1>&2
        exit 1
    fi

    aws eks --region ${REGION} update-kubeconfig --name  kubernetes-${ENVIRONMENT}
    exit_code="$?"
    if [ "$exit_code" != "0" ]; then
        echo "The initialization of the kube config file failed with exit code: ${exit_code}"
    else
        echo "Finished initializing kube-config file for the region:${REGION} environment:${ENVIRONMENT}"
    fi

fi

# Configure it
export CONFIG_FILE="/etc/config.yaml"
export EXEC_CMD="/bin/statusbay -mode=${MODE}"

# Run it:
consul-template \
    -log-level debug \
    -consul-addr=${CONSUL_ADDR} \
    -vault-addr=${VAULT_ADDR} \
    -vault-token=${VAULT_TOKEN} \
    -template "${MODE}.tmpl:${CONFIG_FILE}" \
    -wait 5s \
    -exec-splay 5s \
    -exec "${EXEC_CMD}"
