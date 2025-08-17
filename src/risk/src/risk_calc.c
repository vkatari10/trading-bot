#include <math.h>
#include <stdlib.h>
#include <stdio.h>

double max_drawdown(const double *equity, int n) {
    double peak = equity[0];
    double max_dd = 0.0;
    
    for (int i = 1; i < n; i++) {
        if (equity[i] > peak) peak = equity[i];
        double dd = peak - equity[i];
        if (dd > max_dd) max_dd = dd;
    }
    return max_dd;
}

// -------- Volatility (Std Dev of returns) --------
double volatility(const double *equity, int n) {
    if (n < 2) return 0.0;
    
    double mean_return = 0.0;
    double returns[n-1];
    
    for (int i = 1; i < n; i++) {
        returns[i-1] = equity[i] / equity[i-1] - 1.0;
        mean_return += returns[i-1];
    }
    mean_return /= (n-1);

    double var = 0.0;
    for (int i = 0; i < n-1; i++) {
        var += (returns[i] - mean_return) * (returns[i] - mean_return);
    }
    var /= (n-2); // sample std dev
    return sqrt(var);
}

// -------- Sharpe Ratio --------
double sharpe_ratio(const double *equity, int n, double risk_free_rate) {
    if (n < 2) return 0.0;

    double vol = volatility(equity, n);
    if (vol == 0.0) return 0.0;

    double mean_return = 0.0;
    for (int i = 1; i < n; i++) {
        mean_return += equity[i] / equity[i-1] - 1.0;
    }
    mean_return /= (n-1);

    return (mean_return - risk_free_rate) / vol;
}



