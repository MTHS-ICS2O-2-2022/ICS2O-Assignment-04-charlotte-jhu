// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: April 2023
// This file contains the JS functions for index.html

'use strict'

function myButtonClicked() {
  // input
  const braceletLength = parseFloat(document.getElementById("bracelet-length").getAttribute('data-val'))
  const numberOfColours = parseFloat(document.getElementById("number-of-colours").value)
  const TAX = 0.13

  // process
    const subtotal = braceletLength
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
