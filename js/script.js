// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: April 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
  const TAX = 0.13
  let lengthCost = 0.00
  let colourCost = 0.00

  // input
  const braceletLength = parseFloat(document.getElementById("bracelet-length").value)
  const numberOfColours = parseFloat(document.getElementById("number-of-colours").value)

  // process
  if (braceletLength == "5 inches") {
    lengthCost = 4.00
  } else if (braceletLength == "6 inches") {
    //
  } else if (braceletLength == "7 inches") {
    //
  }

  if (numberOfColours == "1 colour") {
    colourCost = 0.50
  } else if (numberOfColours == "2 colours") {
    //
  } else if (numberOfColours == "3 colours") {
    //
  }


  let subtotal = lengthCost + colourCost



  
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
