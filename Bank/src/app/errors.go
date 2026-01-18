package app

import "errors"

var insufficientFunds error = errors.New("Not enough funds in the account")
var NetworkError error = errors.New("Network error, please try again later")
var WhatTheFuck error = errors.New("What the fuck?")
