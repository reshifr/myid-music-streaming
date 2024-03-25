#!/bin/bash

EXIT_SUCCESS=0
EXIT_FAILURE=1

print_info() { # ( msg... )
	for arg in "$@"; do
		printf "$arg\n"
	done
}

export PGPASSWORD=$POSTGRES_PASSWORD

SQL="\
CREATE ROLE $POSTGRES_USER WITH
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	LOGIN
	REPLICATION
	NOBYPASSRLS
	ENCRYPTED PASSWORD '$POSTGRES_PASSWORD';"

psql \
	--host=$PLAY_POSTGRESQL_IP_ADDRESS0 \
	--username=postgres \
	-c "$SQL"

psql \
	--host=$PLAY_POSTGRESQL_IP_ADDRESS0 \
	--username=postgres \
	-c "DROP ROLE admin;"

psql \
	--host=$PLAY_POSTGRESQL_IP_ADDRESS0 \
	--username=$POSTGRES_USER \
	-c "DROP ROLE postgres;"

psql \
	--host=$PLAY_POSTGRESQL_IP_ADDRESS0 \
	--username=$POSTGRES_USER \
	-c "DROP ROLE postgres;"