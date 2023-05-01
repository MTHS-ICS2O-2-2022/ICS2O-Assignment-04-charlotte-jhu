// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: April 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
  // This function calculates the price of a bracelet
  const TAX = 0.13
  let lengthCost = 0.00
  let colorCost = 0.00

  // input
  const braceletLength = document.getElementById("bracelet-length").value
  const numberOfColors = document.getElementById("number-of-colors").value

  // process
  if (braceletLength == "5 inches") {
    lengthCost = 4.00
  } else if (braceletLength == "6 inches") {
    lengthCost = 4.50
  } else if (braceletLength == "7 inches") {
    lengthCost = 5.00
  }

  if (numberOfColors == "1 colour") {
    colorCost = 0.50
  } else if (numberOfColors == "2 colours") {
    colorCost = 1.00
  } else if (numberOfColors == "3 colours") {
    colorCost = 2.00
  }

  let subtotal = lengthCost + colorCost

  
  if (subtotal >= 6) {
    document.getElementById("special-message").innerHTML = ("Thank you for your purchase, you get a 10% discount next time!")
  }
  const tax = subtotal * TAX
  const total = subtotal * (1 + TAX)

  // output
  document.getElementById("price").innerHTML =
    "Subtotal: $" +
    subtotal.toFixed(2) +
    "<br>Tax: $" +
    tax.toFixed(2) +
    "<br>Total: $" +
    total.toFixed(2)
}
