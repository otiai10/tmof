tmof
=====

How to fight with `too many open files`

steps
=====

1. `ulimit -n`
2. `cat /proc/sys/fs/file-max`
3. `ps aux | grep ${your_app_name}`
4. `sudo ls -l /proc/${proc_id}/fd/ | awk '{print $11}' | awk 'NF' | sort | uniq -c | sort -r`
5. convert hex formatted address to IPv4 string
6. nslookup
