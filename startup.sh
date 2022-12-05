#!/bin/bash
wget https://releases.ubuntu.com/22.04/ubuntu-22.04.1-desktop-amd64.iso
qemu-img create -f qcow2 ubuntu.img 4G
qemu-system-x86_64 -m 256 -hda ubuntu.img -cdrom ubuntu-22.04.1-desktop-amd64.iso -fda boot.bin -nographic

