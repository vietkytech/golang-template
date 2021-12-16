#!/bin/bash
source .env
echo 'project package' ${OLD_PACKAGE} '->' ${CURRENT_PACKAGE}
# sed "s/git.chotot.org\/fse\/multi-rejected-reasons\/multi-rejected-reasons/${PROJECT_PACKAGE}/g" ./apps/
find ./${OLD_APP_NAME} -type f -name "*.go" -exec sed -i "s|${OLD_PACKAGE}|${CURRENT_PACKAGE}|g" {} \;
find ./${OLD_APP_NAME} -type f -name "*.go" -exec sed -i "s|${OLD_APP_NAME}|${APP_NAME}|g" {} \;
# sed -i "s|${OLD_PACKAGE}/${OLD_APP_NAME}|${CURRENT_PACKAGE}/${APP_NAME}|g" ${OLD_APP_NAME}
sed -i "s|${OLD_PACKAGE}|${CURRENT_PACKAGE}|g" Dockerfile
sed -i "s|${OLD_APP_NAME}|${APP_NAME}|g" Dockerfile
# mv `pwd`/multi-rejected-reasons/ apps/
# mv apps/ ./