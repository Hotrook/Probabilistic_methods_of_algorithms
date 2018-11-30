import numpy as np
import matplotlib.pyplot as plt
import argparse

parser = argparse.ArgumentParser(description='Generate chart for permutation experiment.')

parser.add_argument('-datafile', metavar='datafile', 
    help='Filename with data', default='data_50_50_1000.csv')

args = parser.parse_args()

feature_name = ["Fixed points", "Cycles", "Records"]
feature_param = ["Average", "Concentration"]

def read_data_from_file(filename):
    data = np.genfromtxt(filename, delimiter=';')
    (_, sizeY) = np.shape(data)
    return data[:,0], data[:,1:sizeY]

# TODO: add possibility for saving chart to file
def generateChart(x, y):
    for fn_index, fn in enumerate(feature_name):
        for fp_index, fp in enumerate(feature_param):
            plt.subplot(321+fn_index*2+fp_index)
            plt.title(fn + ": " + fp)
            plt.plot(x,y[:,fn_index*2 + fp_index])
    plt.show()

print(args.datafile)
x,y = read_data_from_file(args.datafile)
generateChart(x,y)

