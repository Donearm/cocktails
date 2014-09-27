#!/usr/bin/env python
# -*- coding: utf-8 -*-
###############################################################################
# Copyright (c) 2014, Gianluca Fiore
#
#    This program is free software: you can redistribute it and/or modify
#    it under the terms of the GNU General Public License as published by
#    the Free Software Foundation, either version 3 of the License, or
#    (at your option) any later version.
#
###############################################################################

__author__ = "Gianluca Fiore"
__date__ = "2014/09/27"
__status__ = "stable"

import sys
import json

"""Pretty printing the cocktails in JSON format"""

def main():
    with open(sys.argv[1]) as f:
        j_data = json.load(f)
        print('####')
        print("Cocktail - ", j_data.get('name'))
        if 'description' in j_data:
            print("Description - ", j_data.get('description'))
        print("Method - ", j_data.get('method'))
        print("Glass - ", j_data.get('glass'))
        if 'garnish' in j_data:
            print("Garnish - ", j_data.get('garnish'))
        if 'variations' in j_data:
            print("Variations - ", j_data.get('variations'))
        print('##')

        for e in j_data.get('ingredients'):
            print("Ingredient - ", e.get('ingredientName'))
            if e.get('parts') is not "":
                if e.get('amountUnits') is not "":
                    print(e.get('parts'), "parts /", e.get('amount'), e.get('amountUnits'))
                else:
                    print(e.get('amount'))
            else:
                if e.get('amountUnits') is not "":
                    print(e.get('amount'), e.get('amountUnits'))
                else:
                    print(e.get('amount'))

        print('####')

if __name__ == '__main__':
    status = main()
    sys.exit(status)
