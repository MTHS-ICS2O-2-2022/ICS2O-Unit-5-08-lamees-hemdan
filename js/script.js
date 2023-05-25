// Copyright (c) 2023 Lamees Hemdan All rights reserved
//
// Created by: Lamees Hemdan
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

function doMathClicked () {
 // This function Divides two numbers and shows the answer

  // input
  let number1 = parseInt(document.getElementById("number-1").value)
  let number2 = parseInt(document.getElementById("number-2").value)

 // process
  let counter = 0
  while (number1 >= number2) {
    number1 = number1 - number2
    counter ++
}

 // output
  if (number1 != 0) {
    document.getElementById('number-entered').innerHTML = "the answer is " + counter + " with a remainder of " + number1
  }
  else {
    document.getElementById('number-entered').innerHTML = "The answer is"+ " = " + counter
  }
}
