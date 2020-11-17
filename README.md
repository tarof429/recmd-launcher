# Package for recmd

## Introduction

This is the packaging for recmd.

# Building

To build recmd, you need go 1.15 and modules.

```bash
$ make clean all
```

# Installing

```bash
$ cd ~
$ mkdir recmd
$ cd recmd
$ unzip ../recmd-launcher.zip
```

# Starting

The `recmd` program needs to be started before it can be used. This is because it is a client-server application. 

```bash
$ ./recmd start
```

# Running

The following shows how to add a command and run it in the foreground. The run command also has `-b` option to run commands in the backgruond.

```bash
$ ./recmd list
HASH    COMMAND    DESCRIPTION    DURATION
$ ./recmd add -c "df -h" -d "Disk usage"
Command successfully added.
$ ./recmd list
HASH               COMMAND    DESCRIPTION    DURATION
68e90d0ffb85606    df -h      Disk usage     -
$ ./recmd run 68e90d0ffb85606
✓ Scheduling commmand
Filesystem      Size  Used Avail Use% Mounted on
dev             7.8G     0  7.8G   0% /dev
run             7.9G  1.3M  7.9G   1% /run
/dev/nvme0n1p3  212G   21G  181G  11% /
tmpfs           7.9G  638M  7.2G   8% /dev/shm
tmpfs           4.0M     0  4.0M   0% /sys/fs/cgroup
tmpfs           7.9G   84K  7.9G   1% /tmp
/dev/nvme0n1p2  488M   61M  392M  14% /boot
/dev/sda1       880G   67G  768G   9% /home
/dev/sdb1       916G   95G  775G  11% /mnt/Backup
tmpfs           1.6G   20K  1.6G   1% /run/user/1000
```

You can also install data packs. This lets you imports a collection of premade commands. Below is a data pack to setup Grafana. You may want to delete and recreate the data directory before importing a data pack.

```bash
$ unzip ~/setup-grafana.zip
Archive: setup-grafana.zip
replace conf/recmd_history.json? [y]es, [n]o, [A]ll, [N]one, [r]ename: A
  inflating: conf/recmd_history.json
  inflating: data/install_grafana.sh
  inflating: data/setup_grafana.sh
  inflating: data/setup_grafana_service.sh
  inflating: data/setup.sh
  inflating: data/grafana-server
  inflating: data/download_grafana.sh
  inflating: data/defaults.ini
  inflating: data/env.sh
$ ./recmd start
2020/11/17 14:05:17 Stopping command
2020/11/17 14:05:18 Starting command
$ ./recmd list
HASH               COMMAND                          DESCRIPTION              DURATION
e1280f539e93268    sh ./download_grafana.sh         Download Grafana         -
14bd99d6bb9af5b    sh ./install_grafana.sh          Install Grafana          -
b1d261290dc969a    sh ./setup_grafana.sh            Setup Grafana            0 second(s)
7a359bb41af6c73    sh ./setup_grafana_service.sh    Setup Grafana service    1 second(s)
```

To interactively run through all the commands in order, you can use the `walk` command. This command will display the content of the script that will be executed and prompt the user whether to run it or not. If the answer is `N`, the next command will be shown followed by a user prompt, until all commands have been displayed.

```bash
$ ./recmd walk
This operation will walk through all the commands.
The command will only be run if you answer 'y'.
Do you want to continue? (y/N) y
*****************************
Walking through command 1/4
*****************************

#!/bin/sh

# Load global variables
. ./env.sh

# Step 1: Download the source using curl, untar it, and rename the extracted folder to prometheus-files.
echo "Downloading Prometheus to $TMPDIR"

rm -rf $TMPDIR

mkdir -p $TMPDIR

cd $TMPDIR

wget https://dl.grafana.com/oss/release/grafana-${GRAFANA_VERSION}.linux-amd64.tar.gz
tar -zxvf grafana-${GRAFANA_VERSION}.linux-amd64.tar.gz


Do you want to run this command? (y/N)
```