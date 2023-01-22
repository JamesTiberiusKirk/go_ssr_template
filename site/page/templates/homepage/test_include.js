function render() {
  return elem(
    html.DIV,
    elem(
      html.H1,
      "H1 FROM JS",
    ),
    {
      "style": {
        "border": "blue 2px solid",
      },
    }
  )
}
