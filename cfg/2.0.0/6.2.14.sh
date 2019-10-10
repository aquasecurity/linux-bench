#!/bin/bash 

grep -E -v '^(root|halt|sync|shutdown)' /etc/passwd | awk -F: '($7 != 
"'"$(which nologin)"'" && $7 != "/bin/false") { print $1 " " $6 }' | while 
read user dir; do 
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