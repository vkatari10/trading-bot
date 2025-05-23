import live_recompute as lr
import numpy as np


test = [1, 2, 3, 4, 5, 6]
np_arr = np.array(test)

print(np_arr)
np_stdev = lr.std_dev(np_arr)

print(np_stdev)
print(np_arr)
