#ifndef RISK_CALC_H
#define RISK_CALC_H


// max_drawdown finds the max drawdown in a given input
// 
// args
// arr pointer to the first value in the input
// length length of the array 
// 
// return max_drawdown
double max_drawdown(const double * equity, int n);

double volatility(const double * equity, int n);

double sharpe_ratio(const double * equity, int n, double risk_free_rate);




#endif // RISK_CALC_H