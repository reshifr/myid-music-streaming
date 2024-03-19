#!/bin/bash

EXIT_SUCCESS=0
EXIT_FAILURE=1

exit_success() { # ( msg... )
	for arg in "$@"; do
		printf "$arg\n"
	done
	exit $EXIT_SUCCESS
}

exit_failure() { # ( msg... )
	for arg in "$@"; do
		printf "$arg\n" >&2
	done
	exit $EXIT_FAILURE
}

print_info() { # ( msg... )
	for arg in "$@"; do
		printf "$arg\n"
	done
}

verbosify() { # ( cmd, args... )
	if [ $VERBOSIFY = 0 ]; then
		"$@" &> /dev/null
	else
		"$@"
	fi
}
