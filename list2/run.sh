#!/usr/bin/env bash

N=50
STEP=20
START=20

go install ./...
runExperiment -n=${N} -step=${STEP} -start=${START}
python3 chart.py
