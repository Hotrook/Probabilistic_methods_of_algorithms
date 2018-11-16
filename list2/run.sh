#!/usr/bin/env bash

# sysinfo_page - A script to generate experiment on Random Graphs and show charts from the results.

##### Default constants

# TODO: Add manual for that

DEFAULT_START=50
DEFAULT_STEP=50
DEFAULT_N=500

##### Main

START=${1:-$DEFAULT_START}
STEP=${2:-$DEFAULT_STEP}
N=${3:-$DEFAULT_N}
echo Running experiment with parameters: $'\n'\
    START: ${START} $'\n'\
    STEP: ${STEP} $'\n'\
    N: ${N}

go install ./...
runExperiment -n=${N} -step=${STEP} -start=${START}
python3 charts.py -fst firstComponent_${START}_${STEP}_${N}.txt -snd secondComponent_${START}_${STEP}_${N}.txt
