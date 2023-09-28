import numpy as np

a = np.array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])
b = np.array([0, 2, 4, 6, 8])

print(np.setdiff1d(a, b))

print(np.setdiff1d(b, a))

print(np.intersect1d(a, b))

print(np.union1d(a, b))
