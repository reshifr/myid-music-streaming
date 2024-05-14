#!/bin/bash

EXIT_SUCCESS=0
EXIT_FAILURE=1
MIGRATE_ROLE_DB=migrate_role

info() { # ( msg )
	echo "[INFO] $1"
}

error() { # ( msg )
	echo "[ERROR] $1" >&2
}

catch() { # ( msg )
	if [ $? == $EXIT_SUCCESS ]; then
		info "$1"
	else
		error "$1"
		exit $?
	fi
}

db_wait() {
	while :; do
		if pg_isready \
			--host="$PLAY_POSTGRESQL_IP_ADDRESS0" \
			--username="$PLAY_ADMIN_POSTGRESQL_USER" \
			--dbname="$PLAY_ADMIN_POSTGRESQL_DB" \
      --quiet
		then
			break
		fi
	done
}

admin_sqlexec() { # ( sql )
  PGPASSWORD=$POSTGRES_PASSWORD psql \
    --host="$PLAY_POSTGRESQL_IP_ADDRESS0" \
    --username="$PLAY_ADMIN_POSTGRESQL_USER" \
    --dbname="$PLAY_ADMIN_POSTGRESQL_DB" \
    --no-password \
    --command "$1" 2> /dev/null
}

schema_checking() {
  local sql="\
    CREATE TABLE IF NOT EXISTS $MIGRATE_ROLE_DB (
      version serial PRIMARY KEY,
      timestamp DEFAULT current_timestamp
    )"
  admin_sqlexec "$sql"
}

role_create() {
  case "$1" in
    create)
    ;;
    *)
    ;;
  esac
}

role() {
  case "$1" in
    create) role_create "${@:2}";;
    *) nothing;;
  esac
}

# Call
case "$1" in
  role) role "${@:2}";;
  *) nothing;;
esac
