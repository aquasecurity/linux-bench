#!/bin/bash 

cut -f1 -d":" /etc/group | sort -n | uniq -c | while read x ; do 
	[ -z "$x" ] && break 
	set - $x 
	if [ $1 -gt 1 ]; then 
		gids=$(gawk -F: '($1 == n) { print $3 }' n=$2 /etc/group | xargs) 
		echo "Duplicate Group Name ($2): $gids" 
	fi 
done