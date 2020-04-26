#!/bin/bash

# Set work directory
if [ -n "${WORK_DIR}" ]; then
  cd $WORK_DIR
fi

if [ -z "$GITHUB_TOKEN" ]; then
	echo "Set the GITHUB_TOKEN env variable."
	exit 1
fi

if [ -z "$GITHUB_REPOSITORY" ]; then
	echo "Set the GITHUB_REPOSITORY env variable."
	exit 1
fi

# config
URI=https://api.github.com
API_VERSION=v3
API_HEADER="Accept: application/vnd.github.${API_VERSION}+json; application/vnd.github.antiope-preview+json"
AUTH_HEADER="Authorization: token ${GITHUB_TOKEN}"
TITLE="PHP Security Report(`date "+%Y/%m/%d"`)"

PHP_AUDIT_MESSAG=$(php /opt/security-checker/security-checker security:check composer.lock --format=simple)
PAYLOAD="`/opt/message "${PHP_AUDIT_MESSAG}"`"

# exits error
if [ $? != 0 ]; then
  echo $PAYLOAD
  exit 1
fi

# LABEL setting
ARR=(${ISSUE_LABELS//,/ })
LABEL_ARRAY="["
for STR in "${ARR[@]}";
do
 LABEL_ARRAY="${LABEL_ARRAY}\"${STR}\","
done
LABEL_ARRAY="${LABEL_ARRAY%,}]"

JSON='{"title":"'"${TITLE}"'","body":"'"${PAYLOAD}"'","labels":'${LABEL_ARRAY}'}'

curl -sSL -H "${AUTH_HEADER}" -H "${API_HEADER}" -d "${JSON}" -H "Content-Type: application/json" -X POST "${URI}/repos/${GITHUB_REPOSITORY}/issues"
