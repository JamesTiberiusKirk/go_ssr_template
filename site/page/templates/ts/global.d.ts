export{}

declare interface Data {
}

declare interface Auth {
  email?: string 
  username?: string 
}

declare interface Meta {
  MenuID: string 
  Title: string 
  UrlError: string 
  Success: string 
}


declare interface Routes {
  site: Map<string, string>
}

declare enum Elems  {
  H1     =  "h1",
  H2     =  "h2",
  H3     =  "h3",
  H4     =  "h4",
  H5     =  "h5",
  H6     =  "h6",
  H7     =  "h7",
  DIV    =  "div",
  P      =  "p",
  A      =  "a",
  FORM   =  "form",
  INPUT  =  "input",
  BUTTON =  "button",
}

declare global {
  let elems: Elems
  let _data: Data
  let _auth: Auth
  let _meta: Meta
  let _rotues: Routes

  // function renderFunc(_data:Data): HTMLElement
  function elem(elem: Elems, inner: HTMLElement | string | number): HTMLElement
  function innerTextById(id: string, text: string): void
}

