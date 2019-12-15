#!/bin/sh -l
# Set work directory
if [ -n "${WORK_DIR}" ]; then
  cd $WORK_DIR
fi

PHP_AUDIT_MESSAG=$(php /opt/security-checker/security-checker security:check composer.lock)
STATUS=$?
echo ::set-output name=message::${PHP_AUDIT_MESSAG}
#exit $STATUS