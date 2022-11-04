#!/bin/bash
set -ex

cd /home/zest/Desktop/bridges
sudo chrt -r 77 /bin/bash -c "sudo -u zest go run ."
