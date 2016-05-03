#!/usr/bin/expect -f
set timeout 30
set ipaddress [lindex $argv 0]
spawn ssh yaoyao@relay01.meilishuo.com -A
expect "*"
send "ssh root@${ipaddress}\n"
expect "Last login*" interact
