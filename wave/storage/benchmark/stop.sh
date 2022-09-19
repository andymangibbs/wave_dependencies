#!/bin/bash

TRILLIAN_HOME="/opt/src/github.com/google/trillian"
WAVE_HOME="/opt/src/wave"
# first, stop all trillian server
kill -9 `pgrep trillian`
kill -9 `pgrep "^server"`
