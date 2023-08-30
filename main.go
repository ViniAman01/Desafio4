package main

import (
  ctrl "project/controllers"
)

func main() {
  code_book := "12345"
  ctrl.CoverHandler()
  ctrl.PageHandler(code_book)
}
