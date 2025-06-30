'''
CLI Tool to interact with various services

Author: Vikas Katari 
Date: 06/30/2025
'''
import sys
import src.ml.pipeline as pipe # to train models
import subprocess

args = sys.argv

# TODO implement error handling

if args[1] == "train":
    pipe.pipeline(args[2]) 
elif args[1] == "env":
    pass
elif args[1] == "build":
    pass
elif args[1] == "test":
    subprocess.run(["bash", "./scripts/test.sh"])