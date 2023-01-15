function render() {
  return elem(
    elems.DIV,
    elem(
      elems.H1,
      "H1 FROM JS",
    ),
    {
      "style": {
        "border": "blue 2px solid",
      },
    }
  )
}
