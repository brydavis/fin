from __future__ import division

import math


def future_value(V, R, n, m=1):
    return V * (1+(R/m))**(n*m)

def present_value(FV, R, n):
    return FV / ((1+R)**n)

def compound_return(V, FV, n):
    return ((FV / V)**(1/n))-1

def investment_horizon(V, FV, R):
    return math.log(FV/V)/math.log(1+R)