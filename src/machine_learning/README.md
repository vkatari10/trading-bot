# Machine Learning
This folder contains the entire Machine Learning pipeline from constructing the Pandas DataFrame, training models, and exporting the models.<br>

The feature engineering involves using technical indicators and user defined relationships. The features are declared in `../logic/features.json` as a JSON array.<br>

User defined relationships are defined in the `pipeline.py` file which contains a method to add new columns to the training DataFrame using methods defined in `data_processing/signals.py`.<br>

After defining the wanted technical indicators that are supported and the relationships between the technical indicators, then the `pipeline.py` script can be executed and will train, test, and export the model to the intended folder in `models/decider` that the Machine Learning API will use as th decider in real time inside the runtime environment.<br>


