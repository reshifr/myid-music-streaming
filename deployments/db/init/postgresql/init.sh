#!/bin/bash

EXIT_SUCCESS=0
EXIT_FAILURE=1

info() { # ( msg )
	echo "[INFO] $1"
}

error() { # ( msg )
	echo "[ERROR] $1" >&2
}

catch() { # ( msg, output )
	if [ $? == $EXIT_SUCCESS ]; then
		info "$1"
		echo "$2"
	else
		error "$1"
		echo "$2"
		exit $?
	fi
}

db_wait() {
	while :; do
		info "Trying to connect the database..."
		if pg_isready \
			--host="$PLAY_POSTGRESQL_IP_ADDRESS0" \
			--username="$POSTGRES_USER" \
			--dbname="$POSTGRES_DB"
		then
			break
		fi
	done
}

su_sqlexec() { # ( msg, sql )
	output=$(
		PGPASSWORD=$POSTGRES_PASSWORD psql \
			--host="$PLAY_POSTGRESQL_IP_ADDRESS0" \
			--username="$POSTGRES_USER" \
			--dbname="$POSTGRES_DB" \
			--no-password \
			--command "$2" 2>&1)
	catch "$1" "$output"
}

# Waiting the database availability
db_wait

###################
# Control initial #
###################

# Drop an initialized databases
sql="DROP DATABASE postgres"
su_sqlexec "Drop an initialized databases" "$sql"

# Create an admin role
sql="\
	CREATE ROLE $PLAY_ADMIN_POSTGRESQL_USER WITH
		NOSUPERUSER
		CREATEDB
		CREATEROLE
		INHERIT
		LOGIN
		REPLICATION
		BYPASSRLS
		PASSWORD '$PLAY_ADMIN_POSTGRESQL_PASSWORD'"
su_sqlexec "Create an admin role" "$sql"

# Create an admin database
sql="\
	CREATE DATABASE $PLAY_ADMIN_POSTGRESQL_DB
		WITH OWNER = '$PLAY_ADMIN_POSTGRESQL_USER'"
su_sqlexec "Create an admin database" "$sql"

# Inform the control initial status
info "Control initial succeed"
