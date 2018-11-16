import numpy as np
import matplotlib.pyplot as plt

probabilitiesNumber = 6
names = ["first", "second", "third", "fourth", "fifth", "sixth"]

def read_data_from_file(filename):
    file = open(filename, "r")
    lines = list(file)

    sizeX = len(lines)
    sizeY = probabilitiesNumber 

    x = np.zeros((sizeX))
    y = np.zeros((sizeX, sizeY))

    for index, line in enumerate(lines):
        numbers = line.split(";")
        x[index] = int(numbers[0])
        for p in range(1,sizeY+1):
            y[index][p-1] = float(numbers[p])

    return x, y
        
def generateChart(x, y):
    for p in range(0,probabilitiesNumber):
        plt.subplot(231+p)
        plt.title(names[p])
        plt.plot(x,y[:,p])

    plt.show()

x,y = read_data_from_file("firstComponent.txt")
generateChart(x,y)

x,y = read_data_from_file("secondComponent.txt")
generateChart(x,y)