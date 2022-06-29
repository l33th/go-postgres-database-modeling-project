#!/bin/sh

curl -L $(curl -s https://api.github.com/repos/Masterminds/glide/releases/latest | grep 'browser_' | cut -d\" -f4 | grep 'linux-amd64.tar.gz') > glide.tar.gz &&
tar --strip-components=1 -C /usr/bin -xzf glide.tar.gz linux-amd64/glide &&
rm glide.tar.gz
