#!/bin/bash 

cut -f1 -d":" /etc/passwd | sort -n | uniq -c | while read x ; do 
	[ -z "$x" ] && break 
	set - $x 
	if [ $1 -gt 1 ]; then 
		uids=$(awk -F: '($1 == n) { print $3 }' n=$2 /etc/passwd | xargs) 
		echo "Duplicate User Name ($2): $uids"
	fi 
done