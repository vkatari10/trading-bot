#ifndef RISK_CALC_H
#define RISK_CALC_H

// Author: Vikas Katari
// Date: 08/19/2025

// These 3 methods are AI generated in the risk_calc.c file
// they should be implemented from scratch or at least from more 
// reputable sources before actually using them...

double max_drawdown(const double * equity, int n);

double volatility(const double * equity, int n);

double sharpe_ratio(const double * equity, int n, double risk_free_rate);

#endif // RISK_CALC_H