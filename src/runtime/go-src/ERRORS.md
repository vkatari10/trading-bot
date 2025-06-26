# Error Codes
If a failure occurs during live execution an error code will be printed to `stdout`.  

If you have logging statements enabled you will recieve more detailed information about what error occured.

| Exit Code | Meaning | Quick Troubleshoot | 
|:---------:|---------|-------|
| 0 | Normal Execution | Congrats | 
| 1 | User JSON Config Parse failed | Is the config JSON setup properly? | 
| 2 | ML API Server unreachable | Is the ML API Server Running? |