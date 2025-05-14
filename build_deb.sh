#!/bin/bash

version=0.9-1

mkdir -p yacicctl_$version/usr/local/bin
cp yacicctl yacicctl_$version/usr/local/bin
sleep 5
dpkg-deb  --build yacicctl_$version
