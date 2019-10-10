#!/bin/bash 

grep -E -v '^(root|halt|sync|shutdown)' /etc/passwd | awk -F: '($7 != 
"'"$(which nologin)"'" && $7 != "/bin/false") { print $1 " " $6 }' | while 
read user dir; do 
	if [ ! -d "$dir" ]; then 
		echo "The home directory ($dir) of user $user does not exist." 
	else 
		if [ ! -h "$dir/.netrc" -a -f "$dir/.netrc" ]; then 
			echo ".netrc file $dir/.netrc exists"
		fi 
	fi 
done