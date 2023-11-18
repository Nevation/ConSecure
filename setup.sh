#!/bin/bash

# Setup the tracing
echo "1" > /sys/kernel/debug/tracing/tracing_on
echo "sched:sched_process_exec" > /sys/kernel/debug/tracing/set_event
echo "sched:sched_process_fork" >> /sys/kernel/debug/tracing/set_event
echo "__do_sys_fork" >> /sys/kernel/debug/tracing/set_ftrace_filter 
echo "__do_sys_vfork" >> /sys/kernel/debug/tracing/set_ftrace_filter 

# Setup the trivy
GITHUB_REPO="aquasecurity/trivy"
OS="Linux"
ARCH="64bit"
DOWNLOAD_CMD="curl -LO"
latest_release=$(curl -s "https://api.github.com/repos/$GITHUB_REPO/releases/latest" | grep -o '"tag_name": "v.*"' | cut -d'"' -f4)
download_url="https://github.com/$GITHUB_REPO/releases/download/$latest_release/trivy_${latest_release#v}_${OS}-${ARCH}.tar.gz"
$DOWNLOAD_CMD $download_url

tar xvfzp "trivy_${latest_release#v}_${OS}-${ARCH}.tar.gz" "trivy"
chmod +x trivy
./triviy image redis:latest &> /dev/null