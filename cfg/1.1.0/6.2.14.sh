#!/bin/bash 

cat /etc/passwd | egrep -v '^(root|halt|sync|shutdown)' | awk -F: '($7 != "/sbin/nologin" && $7 != "/bin/false") { print $1 " " $6 }' | while read user dir; do 
	if [ ! -d "$dir" ]; then 
		echo "The home directory ($dir) of user $user does not exist." 
	else 
		for file in $dir/.rhosts; do 
			if [ ! -h "$file" -a -f "$file" ]; then 
				echo ".rhosts file in $dir" 
			fi 
		done 
	fi 
done