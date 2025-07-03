#! /usr/bin/env python3
'''
CLI Tool to interact with various services

Author: Vikas Katari 
Date: 06/30/2025
'''
try:

    import sys # argv
    import subprocess # run shell scripts
    import src.ml.pipeline as pipe # to train models

    args = sys.argv

    # TODO implement error handling

    if args[1] == "train":
        pipe.pipeline(args[2]) 
    elif args[1] == "build":
        subprocess.run(["bash", "./scripts/build.sh"])
    elif args[1] == "test":
        subprocess.run(["bash", "./scripts/test.sh"])
    elif args[1] == "run":
        file = args[2]       

        compile = subprocess.Popen(
            ["go", "build"],
            cwd="./src/runtime/go-src"
        )

        process = subprocess.Popen( # forked
            ["go", "run", ".", file],
            cwd="./src/runtime/go-src"
        )

        process.communicate()
    elif args[1] == "mlapi":
        if len(args) == 3:
            subprocess.run(["bash", "python3", "-m", "src.api.internal.model_api.model_api", args[2]])
        else: 
            print("usage: ./contrade_cli.py mlapi <path_to_config>")
except ModuleNotFoundError:
    print("venv has been not started, run 'source venv/bin/activate'")
except KeyboardInterrupt:
    print('quit')