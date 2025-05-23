import live_recompute as lr
import numpy as np


test = [10, 12, 14, 13, 15]
np_arr = np.array(test)

print(np_arr)
np_stdev = lr.ema(np_arr, 3, 2)

print(np_stdev)
print(np_arr)
