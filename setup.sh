#!/bin/bash

echo "1" > /sys/kernel/debug/tracing/tracing_on
echo "sched:sched_process_exec" > /sys/kernel/debug/tracing/set_event
echo "sched:sched_process_fork" >> /sys/kernel/debug/tracing/set_event
echo "__do_sys_fork" >> /sys/kernel/debug/tracing/set_ftrace_filter 
echo "__do_sys_vfork" >> /sys/kernel/debug/tracing/set_ftrace_filter 
