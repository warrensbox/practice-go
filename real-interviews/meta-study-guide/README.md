## Troubleshooting and Debugging
- Analysing Performance, steps involved, and ending.
- Linux OS performance metrics. Lot of monitoring tools are built on top of these metrics.


### Methodoligies

- provide giudance in choosing the performance tools. Starting point, process and the ending point.
1. AntiPattern: People tend to run commands they know, not trying to understand what the problem and attacking in solving that instead. (Drunk Man Anti Method) Randomly throwing everything at the problem.
2. Maybe network, firewall etc

- Problem Statement Method
    - Why do you think it has a performance problem, is this a new problem or has been there for some time? Something changed recently? Can be expressed in terms of latency, run-time?
    
- Workload Characterization Method
    - Who is the causing the load? Why is the load called? What is the load? How the load changed over time?
    - Solve some issues
 
- USE Method
    - USE : Utilization, Saturation, Error
    - Functional diagram of the system(listing all componenets of the system), 
    and for every resource check **utilization**(busy time), 
    **saturation**(queue length/time), **errors**(easy to interprate).
    - Current tools might not look everywhere, so this method poses question before the ansers, look at place which are sometimes missed.

- CPU Analysis
    - Process get deadlocked/blocked, at some point (paging, context switching, network io)

- CPU Profile Method
    - Flame graph


### Tooling
 
```
Tool Categories : observability, benchmarking, tuning, static performance tuning, profiling, and tracing
Types of resources : CPU, Memory, IO ie Block Devices(disk) and Network Devices
Types of loads : CPU Bound load, Memory bound load, IO bound load
```

#### Observability Tools : Basics

- **uptime**
  - Measure of cpu demand by looking at system(CPU + disks) load averages (no of processes running or are waiting to run)
  - Load average - average no of processes that have to wait for CPU time.
  - *High Level* idea of system usage and how the load changes. 3 numbers -moving load averages at 1, 5, 15 minute.
  - Interpretation : if the load average at 1 min is more than that of 15 min, the load is increasing, or if reverse then load is decreasing. If load is 0.0, then CPU is idle.
  - If load average is greater than no of CPU, meaning more work than what cpu can dispatch. Indicates CPU Saturation
  - Better alternatives : per-CPU utilization - using mpstat -P ALL 1, per-process CPU utilization - top, pidstat

- **dmesg | tail**
  - Lists system messages, errors messages related to performance measures can be looked from here

- **vmstat 1**
  - summary of servers memory utilization statistics, short for virtual memory stat. 
    - **r**: number of process waiting to run on CPU. Value greater than CPU count means saturation of the server.
    - **free** : free memory in kilobytes. Alternative, more elaborate, *free*
    - **si, so**: Page in and page out (paging~swapping). When pages are written from memory into disk, it is pageout. Page in, when data(process data) is brought from disk to memory, in the forms of pages. 
    Pageins are fine, application initialization will have page-ins. Too many page-out indicate that kernel might be spending too much time managing memeory than application processing (thrashing). 
    In case of constant pageouts, check process occupying cpu the most using ps command.
    Might be confused as IO problem, since disk are used as memory, and swapping would require r/w from said device.m
    Swapping can cause a cascading of bad performance since the waiting processes might keep piling on.
    - **us, sy, id, wa, st** : Percent of CPU times : user time(application), system time (kernel and other system processes), idle(cpu is idle, higher the better), wait I/O, and stolen time. 
    First two indicate that the load may be CPU bound load and looking into the processes in pidstat/top might help in identifying the culprit process(es).
    
  -Options - 1: every second t: timestamp column, SM: Data in Megabytes
  

- **mpstat -P ALL 1**
   - CPU time breakdowns per CPU(cores) us ing the -P option. One of the cores/CPU overworking indicate high usage of a single threaded app. 
   - **usr** : percentage of cpu utilization while executing user level application 
   - **sys** : percentage of cpu utilization while executing by kernel

- **pidstat 1**
  - Summary of per process statistics(breakdown), like top, but doesnt clean the screen. 
  Easy to see patterns over time, rolling output.
  - usr, system for each process.  

- **iostat -xz 1**
  - Used for devices (hard disks), to understand the workload applied and resulting performance.
  - Workload metrics
    - **r/s, w/s, rkB/s, wkB/s** : no of reads & writes, no of kB read and written from the attached devices
  - Resulting performance metrics
    - **await** : avg time for io. Time queued or time being serviced or time waiting for the blocked disk .Larger than expected times might mean device saturation.
    - **util** : Percentage of time the device is doing work. Interpretation: more than 60 percent indicate device saturation
  
 
- **free -m**
  - alternate cat /proc/meminfo
  - buffer/cached = sum of buffer and cache. Buffer used for  block device io, cache used by virtual page cache.

- **sar -n DEV 1**
  - Tool to check network throughput and ensure if it is under the limit. rxKbps and txkBps : measure of workload

- **sar -n TCP,ETCP 1**
   - Overview of tcp metrics.
   - **active** and **passive**: outbound and inbound connections. Used as measure of network load on the server
  
- **sar**
   - Statistics archive, for CPU, Memory, IO, Network, stores for a month. -A option for all the records
   - By default, sar command shows stats for a day, but using -s -e specific time period of the day could be used 

- **top**
  - System wide summary. All of above (memory, CPU, IO, network)
  - Consumes cpu to read /proc.
  - % CPU summed across all CPUs.
  
- **ps** 
  - Process status listing
  
 
 
#### Observability Tools : Intermediate

- **strace**:  System call tracer. Translates syscall args. Usful in solving system usage issues.
   - Implementations of strace use ptrace, alternate use perf using perf-trace. Former slows system down 
   so have to be cautious.
   - Blocks the target, slows application down, Shouldnt be used in production.

- **tcpdump**
  - Trace packets. Packets sequences etc. Scalability issue when network is io in high volumes (gigabits). Doesnt scale well.

- **netstat** 
  - Prints network protocol statistics. Different options provide differ information (interface stats, route table etc)
  - Better command (ip table etc) **ss**
- **nicstat**
  - Network interface stats

- **swapon -s**
  - Shows swap device usage
 
- **lsof**
  - Debug tool.  Understand env, who is connected to who. Which files are connected to which process.

- **sar**
  - System activity reporter. Many statistics (TCP, DEV(networking))
  - Complements top by giving statistics from the past. 



#### Network Issues

- ethtool
  - Diagnose network links
  
  
### Sample Examples  

Hint : Get a functional diagram of the environement, makes easier to create a check list.

----


Example 1 : *System is slow*
1. start with command top for processes and cpu
2. Check Disk io (iostat), and network (sar)
What to do in such a case : Quantify the problem, is it latency etc. Check system resources with methodologies, run through the checklist.

Example 2: 
Application Latency is higher.
USE  METHOD:
  1. **top**  command
    - Check cpu summary, process/kernel time, cpu utilization (if it is 100 percent or not).
  2. CPU utilization again with **vmstat** to see paterns. Check memory, if there is enough left and is not leaninig towards saturation point.
  3. **mpstat** to check if maxing out any cpu
  ```
  Utilization and saturation metrics: swapping not too much, enough memory left, cpu are not overloaded, cpu time for kernel/application is not too much, r is not a lot more than cpu present.
  CPU saturation/utiliation is flexible in case of linux,  kernel manages/moves things around, interrups threads etc if needed. same is not the case with io. 
  ```
  4. Check Disk IO utiliation. **iostat**. util column: more than 60 percent utilization might the problem.
  5. Check Network IO utilization **sar -n DEV 1**. 
  6. pidstat for process wise usage of.