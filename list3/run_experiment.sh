#!/usr/bin/env bash

# sysinfo_page - A script to generate experiment on Permutation and show charts from the results.
# Examined feature are:
#   - average number of cycles and concentration
#   - average number of fixed points and concentration
#   - average number of records and concentration
# example usage: ./run.sh 50 50 500
# example usage with default parameters (listed below): ./run.sh

##### Default constants

# TODO: Add manual for that

DEFAULT_START=50
DEFAULT_STEP=50
DEFAULT_N=1000
DEFAULT_PROBES=500

##### Main

START=${1:-$DEFAULT_START}
STEP=${2:-$DEFAULT_STEP}
N=${3:-$DEFAULT_N}
PROBES=${4:-DEFAULT_PROBES}

echo Running experiment with parameters: $'\n'\
    START: ${START} $'\n'\
    STEP: ${STEP} $'\n'\
    N: ${N} $'\n'\
    PROBES: ${PROBES}

go install ./...
permutation -n=${N} -step=${STEP} -start=${START} -probes=${PROBES}
python3 charts.py -datafile data_${START}_${STEP}_${N}.csv
