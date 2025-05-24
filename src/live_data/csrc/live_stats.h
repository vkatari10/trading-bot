#ifndef  LIVE_STATS_H
#define LIVE_STATS_H


/**
 * @brief Finds the mean of an array of doubles
 *
 * @param arr the array containing doubles
 * @param len the length of the array
 * @return the mean of the array
 */
double mean (const double * arr, size_t len);


/**
 * @brief Finds the standard deviation of an array of doubles
 *
 * @param arr the array containing doubles
 * @param len the length of the array
 * @return the standard deviation of the array
 */
double std_dev (const double * arr, size_t len);


#endif // LIVE_STATS_H
