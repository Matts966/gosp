#!/usr/bin/dumb-init /bin/sh
set -e
exec su-exec gosp:gosp $@ 
