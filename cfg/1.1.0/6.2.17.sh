#!/bin/bash 

cat /etc/group | cut -f3 -d":" | sort -n | uniq -c | while read x ; do 
	[ -z "${x}" ] && break 
	set - $x 
	if [ $1 -gt 1 ]; then 
		groups=`awk -F: '($3 == n) { print $1 }' n=$2 /etc/group | xargs` 
		echo "Duplicate GID ($2): ${groups}" 
	fi 
done