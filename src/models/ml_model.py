from sklearn.pipeline import Pipeline
from sklearn.linear_model import LogisticRegression, LinearRegression
from sklearn.preprocessing import StandardScaler
from src.data.yfinance_data import get_data, process_data

df = get_data("AAPL")
process_data("AAPL")

df.dropna()

X = df.iloc[0:5000][['SMA(20)', 'SMA(50)', 'SMA(200)']]
y = (df['Close'].shift(-1) > df['Close']).astype(int)

print(X)
print("buf")
print(y)
