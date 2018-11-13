// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stat_test

import "gonum/mat"

// ASA Car Exposition Data of Ramos and Donoho (1983)
// http://lib.stat.cmu.edu/datasets/cars.desc
// http://lib.stat.cmu.edu/datasets/cars.data
// Columns are: displacement, horsepower, weight, acceleration, MPG.
var carData = mat.NewDense(392, 5, []float64{
	307.0, 130.0, 3504.0, 12.0, 18.0,
	350.0, 165.0, 3693.0, 11.5, 15.0,
	318.0, 150.0, 3436.0, 11.0, 18.0,
	304.0, 150.0, 3433.0, 12.0, 16.0,
	302.0, 140.0, 3449.0, 10.5, 17.0,
	429.0, 198.0, 4341.0, 10.0, 15.0,
	454.0, 220.0, 4354.0, 9.0, 14.0,
	440.0, 215.0, 4312.0, 8.5, 14.0,
	455.0, 225.0, 4425.0, 10.0, 14.0,
	390.0, 190.0, 3850.0, 8.5, 15.0,
	383.0, 170.0, 3563.0, 10.0, 15.0,
	340.0, 160.0, 3609.0, 8.0, 14.0,
	400.0, 150.0, 3761.0, 9.5, 15.0,
	455.0, 225.0, 3086.0, 10.0, 14.0,
	113.0, 95.0, 2372.0, 15.0, 24.0,
	198.0, 95.0, 2833.0, 15.5, 22.0,
	199.0, 97.0, 2774.0, 15.5, 18.0,
	200.0, 85.0, 2587.0, 16.0, 21.0,
	97.0, 88.0, 2130.0, 14.5, 27.0,
	97.0, 46.0, 1835.0, 20.5, 26.0,
	110.0, 87.0, 2672.0, 17.5, 25.0,
	107.0, 90.0, 2430.0, 14.5, 24.0,
	104.0, 95.0, 2375.0, 17.5, 25.0,
	121.0, 113.0, 2234.0, 12.5, 26.0,
	199.0, 90.0, 2648.0, 15.0, 21.0,
	360.0, 215.0, 4615.0, 14.0, 10.0,
	307.0, 200.0, 4376.0, 15.0, 10.0,
	318.0, 210.0, 4382.0, 13.5, 11.0,
	304.0, 193.0, 4732.0, 18.5, 9.0,
	97.0, 88.0, 2130.0, 14.5, 27.0,
	140.0, 90.0, 2264.0, 15.5, 28.0,
	113.0, 95.0, 2228.0, 14.0, 25.0,
	232.0, 100.0, 2634.0, 13.0, 19.0,
	225.0, 105.0, 3439.0, 15.5, 16.0,
	250.0, 100.0, 3329.0, 15.5, 17.0,
	250.0, 88.0, 3302.0, 15.5, 19.0,
	232.0, 100.0, 3288.0, 15.5, 18.0,
	350.0, 165.0, 4209.0, 12.0, 14.0,
	400.0, 175.0, 4464.0, 11.5, 14.0,
	351.0, 153.0, 4154.0, 13.5, 14.0,
	318.0, 150.0, 4096.0, 13.0, 14.0,
	383.0, 180.0, 4955.0, 11.5, 12.0,
	400.0, 170.0, 4746.0, 12.0, 13.0,
	400.0, 175.0, 5140.0, 12.0, 13.0,
	258.0, 110.0, 2962.0, 13.5, 18.0,
	140.0, 72.0, 2408.0, 19.0, 22.0,
	250.0, 100.0, 3282.0, 15.0, 19.0,
	250.0, 88.0, 3139.0, 14.5, 18.0,
	122.0, 86.0, 2220.0, 14.0, 23.0,
	116.0, 90.0, 2123.0, 14.0, 28.0,
	79.0, 70.0, 2074.0, 19.5, 30.0,
	88.0, 76.0, 2065.0, 14.5, 30.0,
	71.0, 65.0, 1773.0, 19.0, 31.0,
	72.0, 69.0, 1613.0, 18.0, 35.0,
	97.0, 60.0, 1834.0, 19.0, 27.0,
	91.0, 70.0, 1955.0, 20.5, 26.0,
	113.0, 95.0, 2278.0, 15.5, 24.0,
	97.5, 80.0, 2126.0, 17.0, 25.0,
	97.0, 54.0, 2254.0, 23.5, 23.0,
	140.0, 90.0, 2408.0, 19.5, 20.0,
	122.0, 86.0, 2226.0, 16.5, 21.0,
	350.0, 165.0, 4274.0, 12.0, 13.0,
	400.0, 175.0, 4385.0, 12.0, 14.0,
	318.0, 150.0, 4135.0, 13.5, 15.0,
	351.0, 153.0, 4129.0, 13.0, 14.0,
	304.0, 150.0, 3672.0, 11.5, 17.0,
	429.0, 208.0, 4633.0, 11.0, 11.0,
	350.0, 155.0, 4502.0, 13.5, 13.0,
	350.0, 160.0, 4456.0, 13.5, 12.0,
	400.0, 190.0, 4422.0, 12.5, 13.0,
	70.0, 97.0, 2330.0, 13.5, 19.0,
	304.0, 150.0, 3892.0, 12.5, 15.0,
	307.0, 130.0, 4098.0, 14.0, 13.0,
	302.0, 140.0, 4294.0, 16.0, 13.0,
	318.0, 150.0, 4077.0, 14.0, 14.0,
	121.0, 112.0, 2933.0, 14.5, 18.0,
	121.0, 76.0, 2511.0, 18.0, 22.0,
	120.0, 87.0, 2979.0, 19.5, 21.0,
	96.0, 69.0, 2189.0, 18.0, 26.0,
	122.0, 86.0, 2395.0, 16.0, 22.0,
	97.0, 92.0, 2288.0, 17.0, 28.0,
	120.0, 97.0, 2506.0, 14.5, 23.0,
	98.0, 80.0, 2164.0, 15.0, 28.0,
	97.0, 88.0, 2100.0, 16.5, 27.0,
	350.0, 175.0, 4100.0, 13.0, 13.0,
	304.0, 150.0, 3672.0, 11.5, 14.0,
	350.0, 145.0, 3988.0, 13.0, 13.0,
	302.0, 137.0, 4042.0, 14.5, 14.0,
	318.0, 150.0, 3777.0, 12.5, 15.0,
	429.0, 198.0, 4952.0, 11.5, 12.0,
	400.0, 150.0, 4464.0, 12.0, 13.0,
	351.0, 158.0, 4363.0, 13.0, 13.0,
	318.0, 150.0, 4237.0, 14.5, 14.0,
	440.0, 215.0, 4735.0, 11.0, 13.0,
	455.0, 225.0, 4951.0, 11.0, 12.0,
	360.0, 175.0, 3821.0, 11.0, 13.0,
	225.0, 105.0, 3121.0, 16.5, 18.0,
	250.0, 100.0, 3278.0, 18.0, 16.0,
	232.0, 100.0, 2945.0, 16.0, 18.0,
	250.0, 88.0, 3021.0, 16.5, 18.0,
	198.0, 95.0, 2904.0, 16.0, 23.0,
	97.0, 46.0, 1950.0, 21.0, 26.0,
	400.0, 150.0, 4997.0, 14.0, 11.0,
	400.0, 167.0, 4906.0, 12.5, 12.0,
	360.0, 170.0, 4654.0, 13.0, 13.0,
	350.0, 180.0, 4499.0, 12.5, 12.0,
	232.0, 100.0, 2789.0, 15.0, 18.0,
	97.0, 88.0, 2279.0, 19.0, 20.0,
	140.0, 72.0, 2401.0, 19.5, 21.0,
	108.0, 94.0, 2379.0, 16.5, 22.0,
	70.0, 90.0, 2124.0, 13.5, 18.0,
	122.0, 85.0, 2310.0, 18.5, 19.0,
	155.0, 107.0, 2472.0, 14.0, 21.0,
	98.0, 90.0, 2265.0, 15.5, 26.0,
	350.0, 145.0, 4082.0, 13.0, 15.0,
	400.0, 230.0, 4278.0, 9.5, 16.0,
	68.0, 49.0, 1867.0, 19.5, 29.0,
	116.0, 75.0, 2158.0, 15.5, 24.0,
	114.0, 91.0, 2582.0, 14.0, 20.0,
	121.0, 112.0, 2868.0, 15.5, 19.0,
	318.0, 150.0, 3399.0, 11.0, 15.0,
	121.0, 110.0, 2660.0, 14.0, 24.0,
	156.0, 122.0, 2807.0, 13.5, 20.0,
	350.0, 180.0, 3664.0, 11.0, 11.0,
	198.0, 95.0, 3102.0, 16.5, 20.0,
	232.0, 100.0, 2901.0, 16.0, 19.0,
	250.0, 100.0, 3336.0, 17.0, 15.0,
	79.0, 67.0, 1950.0, 19.0, 31.0,
	122.0, 80.0, 2451.0, 16.5, 26.0,
	71.0, 65.0, 1836.0, 21.0, 32.0,
	140.0, 75.0, 2542.0, 17.0, 25.0,
	250.0, 100.0, 3781.0, 17.0, 16.0,
	258.0, 110.0, 3632.0, 18.0, 16.0,
	225.0, 105.0, 3613.0, 16.5, 18.0,
	302.0, 140.0, 4141.0, 14.0, 16.0,
	350.0, 150.0, 4699.0, 14.5, 13.0,
	318.0, 150.0, 4457.0, 13.5, 14.0,
	302.0, 140.0, 4638.0, 16.0, 14.0,
	304.0, 150.0, 4257.0, 15.5, 14.0,
	98.0, 83.0, 2219.0, 16.5, 29.0,
	79.0, 67.0, 1963.0, 15.5, 26.0,
	97.0, 78.0, 2300.0, 14.5, 26.0,
	76.0, 52.0, 1649.0, 16.5, 31.0,
	83.0, 61.0, 2003.0, 19.0, 32.0,
	90.0, 75.0, 2125.0, 14.5, 28.0,
	90.0, 75.0, 2108.0, 15.5, 24.0,
	116.0, 75.0, 2246.0, 14.0, 26.0,
	120.0, 97.0, 2489.0, 15.0, 24.0,
	108.0, 93.0, 2391.0, 15.5, 26.0,
	79.0, 67.0, 2000.0, 16.0, 31.0,
	225.0, 95.0, 3264.0, 16.0, 19.0,
	250.0, 105.0, 3459.0, 16.0, 18.0,
	250.0, 72.0, 3432.0, 21.0, 15.0,
	250.0, 72.0, 3158.0, 19.5, 15.0,
	400.0, 170.0, 4668.0, 11.5, 16.0,
	350.0, 145.0, 4440.0, 14.0, 15.0,
	318.0, 150.0, 4498.0, 14.5, 16.0,
	351.0, 148.0, 4657.0, 13.5, 14.0,
	231.0, 110.0, 3907.0, 21.0, 17.0,
	250.0, 105.0, 3897.0, 18.5, 16.0,
	258.0, 110.0, 3730.0, 19.0, 15.0,
	225.0, 95.0, 3785.0, 19.0, 18.0,
	231.0, 110.0, 3039.0, 15.0, 21.0,
	262.0, 110.0, 3221.0, 13.5, 20.0,
	302.0, 129.0, 3169.0, 12.0, 13.0,
	97.0, 75.0, 2171.0, 16.0, 29.0,
	140.0, 83.0, 2639.0, 17.0, 23.0,
	232.0, 100.0, 2914.0, 16.0, 20.0,
	140.0, 78.0, 2592.0, 18.5, 23.0,
	134.0, 96.0, 2702.0, 13.5, 24.0,
	90.0, 71.0, 2223.0, 16.5, 25.0,
	119.0, 97.0, 2545.0, 17.0, 24.0,
	171.0, 97.0, 2984.0, 14.5, 18.0,
	90.0, 70.0, 1937.0, 14.0, 29.0,
	232.0, 90.0, 3211.0, 17.0, 19.0,
	115.0, 95.0, 2694.0, 15.0, 23.0,
	120.0, 88.0, 2957.0, 17.0, 23.0,
	121.0, 98.0, 2945.0, 14.5, 22.0,
	121.0, 115.0, 2671.0, 13.5, 25.0,
	91.0, 53.0, 1795.0, 17.5, 33.0,
	107.0, 86.0, 2464.0, 15.5, 28.0,
	116.0, 81.0, 2220.0, 16.9, 25.0,
	140.0, 92.0, 2572.0, 14.9, 25.0,
	98.0, 79.0, 2255.0, 17.7, 26.0,
	101.0, 83.0, 2202.0, 15.3, 27.0,
	305.0, 140.0, 4215.0, 13.0, 17.5,
	318.0, 150.0, 4190.0, 13.0, 16.0,
	304.0, 120.0, 3962.0, 13.9, 15.5,
	351.0, 152.0, 4215.0, 12.8, 14.5,
	225.0, 100.0, 3233.0, 15.4, 22.0,
	250.0, 105.0, 3353.0, 14.5, 22.0,
	200.0, 81.0, 3012.0, 17.6, 24.0,
	232.0, 90.0, 3085.0, 17.6, 22.5,
	85.0, 52.0, 2035.0, 22.2, 29.0,
	98.0, 60.0, 2164.0, 22.1, 24.5,
	90.0, 70.0, 1937.0, 14.2, 29.0,
	91.0, 53.0, 1795.0, 17.4, 33.0,
	225.0, 100.0, 3651.0, 17.7, 20.0,
	250.0, 78.0, 3574.0, 21.0, 18.0,
	250.0, 110.0, 3645.0, 16.2, 18.5,
	258.0, 95.0, 3193.0, 17.8, 17.5,
	97.0, 71.0, 1825.0, 12.2, 29.5,
	85.0, 70.0, 1990.0, 17.0, 32.0,
	97.0, 75.0, 2155.0, 16.4, 28.0,
	140.0, 72.0, 2565.0, 13.6, 26.5,
	130.0, 102.0, 3150.0, 15.7, 20.0,
	318.0, 150.0, 3940.0, 13.2, 13.0,
	120.0, 88.0, 3270.0, 21.9, 19.0,
	156.0, 108.0, 2930.0, 15.5, 19.0,
	168.0, 120.0, 3820.0, 16.7, 16.5,
	350.0, 180.0, 4380.0, 12.1, 16.5,
	350.0, 145.0, 4055.0, 12.0, 13.0,
	302.0, 130.0, 3870.0, 15.0, 13.0,
	318.0, 150.0, 3755.0, 14.0, 13.0,
	98.0, 68.0, 2045.0, 18.5, 31.5,
	111.0, 80.0, 2155.0, 14.8, 30.0,
	79.0, 58.0, 1825.0, 18.6, 36.0,
	122.0, 96.0, 2300.0, 15.5, 25.5,
	85.0, 70.0, 1945.0, 16.8, 33.5,
	305.0, 145.0, 3880.0, 12.5, 17.5,
	260.0, 110.0, 4060.0, 19.0, 17.0,
	318.0, 145.0, 4140.0, 13.7, 15.5,
	302.0, 130.0, 4295.0, 14.9, 15.0,
	250.0, 110.0, 3520.0, 16.4, 17.5,
	231.0, 105.0, 3425.0, 16.9, 20.5,
	225.0, 100.0, 3630.0, 17.7, 19.0,
	250.0, 98.0, 3525.0, 19.0, 18.5,
	400.0, 180.0, 4220.0, 11.1, 16.0,
	350.0, 170.0, 4165.0, 11.4, 15.5,
	400.0, 190.0, 4325.0, 12.2, 15.5,
	351.0, 149.0, 4335.0, 14.5, 16.0,
	97.0, 78.0, 1940.0, 14.5, 29.0,
	151.0, 88.0, 2740.0, 16.0, 24.5,
	97.0, 75.0, 2265.0, 18.2, 26.0,
	140.0, 89.0, 2755.0, 15.8, 25.5,
	98.0, 63.0, 2051.0, 17.0, 30.5,
	98.0, 83.0, 2075.0, 15.9, 33.5,
	97.0, 67.0, 1985.0, 16.4, 30.0,
	97.0, 78.0, 2190.0, 14.1, 30.5,
	146.0, 97.0, 2815.0, 14.5, 22.0,
	121.0, 110.0, 2600.0, 12.8, 21.5,
	80.0, 110.0, 2720.0, 13.5, 21.5,
	90.0, 48.0, 1985.0, 21.5, 43.1,
	98.0, 66.0, 1800.0, 14.4, 36.1,
	78.0, 52.0, 1985.0, 19.4, 32.8,
	85.0, 70.0, 2070.0, 18.6, 39.4,
	91.0, 60.0, 1800.0, 16.4, 36.1,
	260.0, 110.0, 3365.0, 15.5, 19.9,
	318.0, 140.0, 3735.0, 13.2, 19.4,
	302.0, 139.0, 3570.0, 12.8, 20.2,
	231.0, 105.0, 3535.0, 19.2, 19.2,
	200.0, 95.0, 3155.0, 18.2, 20.5,
	200.0, 85.0, 2965.0, 15.8, 20.2,
	140.0, 88.0, 2720.0, 15.4, 25.1,
	225.0, 100.0, 3430.0, 17.2, 20.5,
	232.0, 90.0, 3210.0, 17.2, 19.4,
	231.0, 105.0, 3380.0, 15.8, 20.6,
	200.0, 85.0, 3070.0, 16.7, 20.8,
	225.0, 110.0, 3620.0, 18.7, 18.6,
	258.0, 120.0, 3410.0, 15.1, 18.1,
	305.0, 145.0, 3425.0, 13.2, 19.2,
	231.0, 165.0, 3445.0, 13.4, 17.7,
	302.0, 139.0, 3205.0, 11.2, 18.1,
	318.0, 140.0, 4080.0, 13.7, 17.5,
	98.0, 68.0, 2155.0, 16.5, 30.0,
	134.0, 95.0, 2560.0, 14.2, 27.5,
	119.0, 97.0, 2300.0, 14.7, 27.2,
	105.0, 75.0, 2230.0, 14.5, 30.9,
	134.0, 95.0, 2515.0, 14.8, 21.1,
	156.0, 105.0, 2745.0, 16.7, 23.2,
	151.0, 85.0, 2855.0, 17.6, 23.8,
	119.0, 97.0, 2405.0, 14.9, 23.9,
	131.0, 103.0, 2830.0, 15.9, 20.3,
	163.0, 125.0, 3140.0, 13.6, 17.0,
	121.0, 115.0, 2795.0, 15.7, 21.6,
	163.0, 133.0, 3410.0, 15.8, 16.2,
	89.0, 71.0, 1990.0, 14.9, 31.5,
	98.0, 68.0, 2135.0, 16.6, 29.5,
	231.0, 115.0, 3245.0, 15.4, 21.5,
	200.0, 85.0, 2990.0, 18.2, 19.8,
	140.0, 88.0, 2890.0, 17.3, 22.3,
	232.0, 90.0, 3265.0, 18.2, 20.2,
	225.0, 110.0, 3360.0, 16.6, 20.6,
	305.0, 130.0, 3840.0, 15.4, 17.0,
	302.0, 129.0, 3725.0, 13.4, 17.6,
	351.0, 138.0, 3955.0, 13.2, 16.5,
	318.0, 135.0, 3830.0, 15.2, 18.2,
	350.0, 155.0, 4360.0, 14.9, 16.9,
	351.0, 142.0, 4054.0, 14.3, 15.5,
	267.0, 125.0, 3605.0, 15.0, 19.2,
	360.0, 150.0, 3940.0, 13.0, 18.5,
	89.0, 71.0, 1925.0, 14.0, 31.9,
	86.0, 65.0, 1975.0, 15.2, 34.1,
	98.0, 80.0, 1915.0, 14.4, 35.7,
	121.0, 80.0, 2670.0, 15.0, 27.4,
	183.0, 77.0, 3530.0, 20.1, 25.4,
	350.0, 125.0, 3900.0, 17.4, 23.0,
	141.0, 71.0, 3190.0, 24.8, 27.2,
	260.0, 90.0, 3420.0, 22.2, 23.9,
	105.0, 70.0, 2200.0, 13.2, 34.2,
	105.0, 70.0, 2150.0, 14.9, 34.5,
	85.0, 65.0, 2020.0, 19.2, 31.8,
	91.0, 69.0, 2130.0, 14.7, 37.3,
	151.0, 90.0, 2670.0, 16.0, 28.4,
	173.0, 115.0, 2595.0, 11.3, 28.8,
	173.0, 115.0, 2700.0, 12.9, 26.8,
	151.0, 90.0, 2556.0, 13.2, 33.5,
	98.0, 76.0, 2144.0, 14.7, 41.5,
	89.0, 60.0, 1968.0, 18.8, 38.1,
	98.0, 70.0, 2120.0, 15.5, 32.1,
	86.0, 65.0, 2019.0, 16.4, 37.2,
	151.0, 90.0, 2678.0, 16.5, 28.0,
	140.0, 88.0, 2870.0, 18.1, 26.4,
	151.0, 90.0, 3003.0, 20.1, 24.3,
	225.0, 90.0, 3381.0, 18.7, 19.1,
	97.0, 78.0, 2188.0, 15.8, 34.3,
	134.0, 90.0, 2711.0, 15.5, 29.8,
	120.0, 75.0, 2542.0, 17.5, 31.3,
	119.0, 92.0, 2434.0, 15.0, 37.0,
	108.0, 75.0, 2265.0, 15.2, 32.2,
	86.0, 65.0, 2110.0, 17.9, 46.6,
	156.0, 105.0, 2800.0, 14.4, 27.9,
	85.0, 65.0, 2110.0, 19.2, 40.8,
	90.0, 48.0, 2085.0, 21.7, 44.3,
	90.0, 48.0, 2335.0, 23.7, 43.4,
	121.0, 67.0, 2950.0, 19.9, 36.4,
	146.0, 67.0, 3250.0, 21.8, 30.0,
	91.0, 67.0, 1850.0, 13.8, 44.6,
	97.0, 67.0, 2145.0, 18.0, 33.8,
	89.0, 62.0, 1845.0, 15.3, 29.8,
	168.0, 132.0, 2910.0, 11.4, 32.7,
	70.0, 100.0, 2420.0, 12.5, 23.7,
	122.0, 88.0, 2500.0, 15.1, 35.0,
	107.0, 72.0, 2290.0, 17.0, 32.4,
	135.0, 84.0, 2490.0, 15.7, 27.2,
	151.0, 84.0, 2635.0, 16.4, 26.6,
	156.0, 92.0, 2620.0, 14.4, 25.8,
	173.0, 110.0, 2725.0, 12.6, 23.5,
	135.0, 84.0, 2385.0, 12.9, 30.0,
	79.0, 58.0, 1755.0, 16.9, 39.1,
	86.0, 64.0, 1875.0, 16.4, 39.0,
	81.0, 60.0, 1760.0, 16.1, 35.1,
	97.0, 67.0, 2065.0, 17.8, 32.3,
	85.0, 65.0, 1975.0, 19.4, 37.0,
	89.0, 62.0, 2050.0, 17.3, 37.7,
	91.0, 68.0, 1985.0, 16.0, 34.1,
	105.0, 63.0, 2215.0, 14.9, 34.7,
	98.0, 65.0, 2045.0, 16.2, 34.4,
	98.0, 65.0, 2380.0, 20.7, 29.9,
	105.0, 74.0, 2190.0, 14.2, 33.0,
	107.0, 75.0, 2210.0, 14.4, 33.7,
	108.0, 75.0, 2350.0, 16.8, 32.4,
	119.0, 100.0, 2615.0, 14.8, 32.9,
	120.0, 74.0, 2635.0, 18.3, 31.6,
	141.0, 80.0, 3230.0, 20.4, 28.1,
	145.0, 76.0, 3160.0, 19.6, 30.7,
	168.0, 116.0, 2900.0, 12.6, 25.4,
	146.0, 120.0, 2930.0, 13.8, 24.2,
	231.0, 110.0, 3415.0, 15.8, 22.4,
	350.0, 105.0, 3725.0, 19.0, 26.6,
	200.0, 88.0, 3060.0, 17.1, 20.2,
	225.0, 85.0, 3465.0, 16.6, 17.6,
	112.0, 88.0, 2605.0, 19.6, 28.0,
	112.0, 88.0, 2640.0, 18.6, 27.0,
	112.0, 88.0, 2395.0, 18.0, 34.0,
	112.0, 85.0, 2575.0, 16.2, 31.0,
	135.0, 84.0, 2525.0, 16.0, 29.0,
	151.0, 90.0, 2735.0, 18.0, 27.0,
	140.0, 92.0, 2865.0, 16.4, 24.0,
	105.0, 74.0, 1980.0, 15.3, 36.0,
	91.0, 68.0, 2025.0, 18.2, 37.0,
	91.0, 68.0, 1970.0, 17.6, 31.0,
	105.0, 63.0, 2125.0, 14.7, 38.0,
	98.0, 70.0, 2125.0, 17.3, 36.0,
	120.0, 88.0, 2160.0, 14.5, 36.0,
	107.0, 75.0, 2205.0, 14.5, 36.0,
	108.0, 70.0, 2245.0, 16.9, 34.0,
	91.0, 67.0, 1965.0, 15.0, 38.0,
	91.0, 67.0, 1965.0, 15.7, 32.0,
	91.0, 67.0, 1995.0, 16.2, 38.0,
	181.0, 110.0, 2945.0, 16.4, 25.0,
	262.0, 85.0, 3015.0, 17.0, 38.0,
	156.0, 92.0, 2585.0, 14.5, 26.0,
	232.0, 112.0, 2835.0, 14.7, 22.0,
	144.0, 96.0, 2665.0, 13.9, 32.0,
	135.0, 84.0, 2370.0, 13.0, 36.0,
	151.0, 90.0, 2950.0, 17.3, 27.0,
	140.0, 86.0, 2790.0, 15.6, 27.0,
	97.0, 52.0, 2130.0, 24.6, 44.0,
	135.0, 84.0, 2295.0, 11.6, 32.0,
	120.0, 79.0, 2625.0, 18.6, 28.0,
	119.0, 82.0, 2720.0, 19.4, 31.0,
})
