Restart all instances of a running service, for example with `service elasticsearch restart`, and ensure that no more than N% of the service is restarting at the # same time. 
Connect to each host with ssh.

Hosts are in a hosts.txt file with one hostname per line:
host1 1
host2 0
3.    1
...
hostN

Progression:
- Restart one instance at a time
- Restart two instances at a time
- Restart N instances at a time
- Restart N% of the instances at a time

For each iteration, check the status of the service.

Follow-up:
- Error handling: the goal is to restart as many service instances as possible and to keep a list of services that failed
- Handle when a host is unreachable
- Handle when the command fails on a host (ex: service no running)