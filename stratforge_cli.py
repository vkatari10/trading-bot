
import cli.jsonwriter
import subprocess
import sys

args = sys.argv

def handle_command(args) -> None:

    command = ""

    if len(args) == 1:
        # env, test
        command = f"scripts/{command}.sh"
        subprocess.run(["bash", args])
    elif len(args) == 2:
        # build, mlapi, train, run (soon) need to 
        pass



   
if args[1] == "config":
    cli.jsonwriter.script()

handle_command(args) # run other command

