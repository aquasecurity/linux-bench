#!/bin/bash

grep -E -v '^(halt|sync|shutdown)' /etc/passwd | awk -F: '($7 != "'"$(which 
nologin)"'" && $7 != "/bin/false") { print $1 " " $6 }' | while read user 
dir; do 
	if [ ! -d "$dir" ]; then 
		echo "The home directory ($dir) of user $user does not exist." 
	else 
		owner=$(stat -L -c "%U" "$dir") 
		if [ "$owner" != "$user" ]; then 
			echo "The home directory ($dir) of user $user is owned by $owner."
		fi 
	fi 
done