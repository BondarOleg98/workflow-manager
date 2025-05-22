#!/bin/bash

resources_path="app/resources"
yq=./${resources_path}/yq

curl -L -o ${yq} https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64
chmod +x ${yq}

#${yq} ${resources_path}/env.yaml -os > ${resources_path}/env.sh
#source ${resources_path}/env.sh

#rm ${yq}
#rm ${resources_path}/env.sh
#
#chmod +x ./app/db/postgres/initdb.sh
#./app/db/postgres/initdb.sh