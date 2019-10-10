#!/bin/bash 

cut -f3 -d":" /etc/passwd | sort -n | uniq -c | while read x ; do 
	[ -z "$x" ] && break 
	set - $x 
	if [ $1 -gt 1 ]; then 
		users=$(awk -F: '($3 == n) { print $1 }' n=$2 /etc/passwd | xargs) 
		echo "Duplicate UID ($2): $users"
	fi 
done