#!/bin/bash

logInfo "Run the db migrations script"
# shellcheck disable=SC2154
runSQL "${POSTGRES_DB}" "-f" "${pg_scripts_path}/${POSTGRES_MIGRATIONS_SCRIPT}"