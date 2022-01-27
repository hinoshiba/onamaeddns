#!/bin/bash
set -eu

OLD_PATH="/tmp/.ddns_clnt"
HOST="${TARGET_HOST}"
DOMAIN="${TARGET_DOMAIN}"

echo "Start onamaeddns. target: '${HOST}.${DOMAIN}'"

while :
do
	test -f "${OLD_PATH}" || :>"${OLD_PATH}"
	n_gip=$(curl -s globalip.me)
	o_gip=$(cat ${OLD_PATH})

	if [ "${n_gip}" != "${o_gip}" ]; then
		echo "Detect new global ip address. ${n_gip}"
		echo -n "exec onamaeddns....."
		onamaeddns -c /etc/onamaeddns/cred ${HOST} ${DOMAIN} ${n_gip}
		echo "done"
		echo -n ${n_gip} > ${OLD_PATH}
	fi

	sleep 300
done
