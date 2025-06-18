import src.ml.data_processing.signals as sig
import pandas as pd

def get_dummy_stock_df() -> pd.DataFrame:
    data = {
        "Date": pd.date_range(start="2024-01-02", periods=5, freq="D"),
        "Open": [150.00, 153.50, 157.25, 158.75, 161.00],
        "High": [155.00, 158.00, 160.00, 162.00, 163.00],
        "Low":  [149.00, 152.00, 155.00, 157.00, 159.00],
        "Close": [153.50, 157.25, 155.75, 161.00, 160.25],
        "Adj Close": [153.50, 157.25, 158.75, 161.00, 160.25],
        "Volume": [1_000_000, 1_200_000, 1_100_000, 1_050_000, 980_000]
    }
    return pd.DataFrame(data)


def test_signals():
    '''
    Test if signals come out as intended on a dummy DataFrame
    '''
    test_df = get_dummy_stock_df()

    test_df['Crossover'] = sig.crossover(test_df, "Open", "Close")
    test_df['Above'] = sig.above(test_df, "Open", "Close")
    test_df['Below'] = sig.below(test_df, "Open", "Close")

    assert test_df['Crossover'].equals(pd.Series([0, 0, 1, -1, 1]))
    assert test_df['Above'].equals(pd.Series([0, 0, 1, 0, 1]))
    assert test_df['Below'].equals(pd.Series([1, 1, 0, 1, 0]))