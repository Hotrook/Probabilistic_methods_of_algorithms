#!/usr/bin/env bash

# sysinfo_page - A script to generate experiment on Random Graphs and show charts from the results.
# example usage: ./run.sh 50 50 500 100
# example usage with default parameters (listed below): ./run.sh

##### Default constants

# TODO: Add manual for that

DEFAULT_START=50
DEFAULT_STEP=50
DEFAULT_N=500
DEFAULT_PROBES=100

##### Main

START=${1:-$DEFAULT_START}
STEP=${2:-$DEFAULT_STEP}
N=${3:-$DEFAULT_N}
PROBES=${4:-$DEFAULT_PROBES}

echo Running experiment with parameters: $'\n'\
    START: ${START} $'\n'\
    STEP: ${STEP} $'\n'\
    N: ${N} $'\n'\
    PROBES: ${PROBES}

mkdir results
go install ./...
runExperiment -n=${N} -step=${STEP} -start=${START} -probes=${PROBES}
python3 charts.py -fst results/firstComponent_${START}_${STEP}_${N}.txt -snd results/secondComponent_${START}_${STEP}_${N}.txt
