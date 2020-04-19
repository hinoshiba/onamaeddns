#!/bin/bash

OLD_PATH="${HOME}/.ddns_clnt"
HOST="test"
DOMAIN="example.com"

n_gip=$(curl -s globalip.me)
o_gip=$(cat ${OLD_PATH})

if [ "${n_gip}" != "${o_gip}" ]; then
	onamaeddns ${HOST} ${DOMAIN} ${n_gip}
	echo -n ${n_gip} > ${OLD_PATH}
fi
