# Error Codes

If a failure occurs during live execution an error code will be printed to `stdout`.  

If you have logging statements enabled you will recieve more detailed information about what error occured.

| Exit Code | Meaning | Quick Troubleshoot | 
|:---------:|---------|-------|
| 0 | Normal Execution | Congrats | 
| 1 | Provided JSON Config Failed to Initialize | Check if your config file has any errors, to learn more see [documentation](CONFIG.md) | 
| 2 | ML API could not be found | Make sure the ML API is running with the selected config | 
| 3 | JSON object representation failed to be initialized | Check if your config file has any errors, to learn more see [documentation](CONFIG.md) | 



<!-- | Exit Code | Meaning | Quick Troubleshoot | 
|:---------:|---------|-------|
| 0 | Normal Execution | Congrats | 
| 1 | JSON 'features' array of objects failed to initialize |  Ensure your features are declared properly, see [docs](CONFIG.md) for more info. | 
| 2 | JSON 'runtime_settings' object could not be read | Ensure your runtime settings are declared properly, see [docs](CONFIG.md) for more info | 
| 3 | JSON 'live_trade_tickers` array could not be read | Ensure there is at least one live ticker declared, see [docs](CONFIG.md) for more info | 
| 4 | ML API Server unreachable | Is the ML API Server Running? | -->
