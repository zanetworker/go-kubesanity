#!/usr/bin/env bash
set -euo pipefail

bumpversion(){
	declare versionPath="$1" versionType="$2"

	if [[ -n "$versionPath" ]]; then 
		version="$(cat $versionPath)"

		echo "Current Version is: '$version'"

		local major="$(cut -d'.' -f1 <<< $version)"
		local minor="$(cut -d'.' -f2 <<< $version)"
		local patch="$(cut -d'.' -f3 <<< $version)"

		local versionArray=("$major" "$minor" "$patch")
		_is_number "$versionArray"

		if [[ -n "$versionPath" ]]; then 
			case "$versionType" in
				major) let "major++" ;;
				minor) let "minor++" ;; 
				patch) let "patch++" ;; 
			esac
		fi 
		
		echo "$major.$minor.$patch" > "$versionPath"
		echo "Bumped version is '$(cat $versionPath)''"
	fi 
}

_is_number(){
	declare versionNumbersToCheck="$@"

    for i in "$versionNumbersToCheck"; do 
		local re='^[0-9]+$'
		if ! [[ "$i" =~ ^-?[0-9]+$ ]]; then echo "error: Not a number" >&2; exit 1; fi
	done
}

_help() {
	echo "You need to supply the right parameters to the script"
	echo "Usage: \"bumpversion _path_to_version_file major|minor|patch\"" 
}

main(){
	if [[ "$#" -eq 0 ]]; then
    	_help 
		exit 1 
	fi
	bumpversion "$1" "$2" 
}


main "$@"