import cv2
import numpy as np
import os
import matplotlib.pyplot as plt
import time

def loadImages(basePath, classes):
    imageVectors = []
    labels = []

    for className in classes:
        classPath = os.path.join(basePath, className)
        for imageName in os.listdir(classPath):
            imagePath = os.path.join(classPath, imageName)
            image = cv2.imread(imagePath)
            if image is not None:
                image = cv2.resize(image, (32, 32))
                _, binaryImage = cv2.threshold(image, 127, 256, cv2.THRESH_BINARY)
                imageVector = binaryImage.flatten()
                imageVectors.append(imageVector)
                labels.append(className)

    return imageVectors, labels

def loadTestImages(basePath):
    imageVectors = []

    for imageName in os.listdir(basePath):
        imagePath = os.path.join(basePath, imageName)
        image = cv2.imread(imagePath)
        if image is not None:
            image = cv2.resize(image, (32, 32))
            _, binaryImage = cv2.threshold(image, 127, 256, cv2.THRESH_BINARY)
            imageVector = binaryImage.flatten()
            imageVectors.append(imageVector)

    return imageVectors

def hammingDistance(vector1, vector2):
    return np.sum(vector1 != vector2)


def printHammingDistance(hD):
    print(f"Hamming Distance: {hD}")

def findPotential(R):
    return 1000000/(1 + R**2)

def getAnswer(classes, trainVectors, trainLabels, testVectors):
    plt.figure(figsize=(8, 8))
    plt.ion()
    for i in range(0, len(testVectors)):
        classPotentials = {}
        for className in classes:
            classPotentials[className] = 0


        for j in range(0, len(trainVectors)):
            hD = hammingDistance(testVectors[i], trainVectors[j])
            potential = findPotential(hD)
            classPotentials[trainLabels[j]] += potential
            
        print(f"Потенциалы для тествого {i}:")
        for className, potential in classPotentials.items():
            print(f"  {className}: {potential:.2f}")
        winner = max(classPotentials, key=classPotentials.get)

        imgFromVector = testVectors[i].reshape(32, 32, 3)

        plt.clf()
        plt.imshow(imgFromVector, cmap='gray')
        plt.title(f'Image: {winner}')
        potentialsText = ' '.join([f'{className}: {potential:.2f}' for className, potential in classPotentials.items()])
        plt.figtext(0.5, 0.03, f'Potentials:\n{potentialsText}', ha='center', fontsize=10, bbox=dict(facecolor='white', alpha=0.5))
        plt.draw()
        plt.pause(5)
    plt.ioff()
    plt.show()



classes = ['circle', 'rect', 'triangle']
trainVectors, trainLabels = loadImages('./Samples', classes)
testVectors = loadTestImages('./Samples/test')
getAnswer(classes, trainVectors, trainLabels, testVectors)



