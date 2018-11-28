import numpy as np
import matplotlib.pyplot as plt
import argparse

parser = argparse.ArgumentParser(description='Generate chart for Random Graphs experiments.')
parser.add_argument('-fst', metavar='fst', 
    help='Filename with data for first component', default="firstComponent.txt")
parser.add_argument('-snd', metavar='snd', 
    help='Filename with data for second component', default="secondComponent.txt")

args = parser.parse_args()

probabilitiesNumber = 6
names = ["first", "second", "third", "fourth", "fifth", "sixth"]

def read_data_from_file(filename):
    data = np.genfromtxt(filename, delimiter=';')
    (_, sizeY) = np.shape(data)
    return data[:,0], data[:,1:sizeY]

# TODO: add possibility for saving chart to file
def generateChart(x, y):
    for p in range(0,probabilitiesNumber):
        plt.subplot(231+p)
        plt.title(names[p])
        plt.plot(x,y[:,p])

    plt.show()

x,y = read_data_from_file(args.fst)
generateChart(x,y)

x,y = read_data_from_file(args.snd)
generateChart(x,y)

