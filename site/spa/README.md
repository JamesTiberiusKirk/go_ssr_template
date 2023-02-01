# SPA
Created with...
```sh
npx create-react-app react_portal --template typescript
```

## To get working
- In order to get the react app to work in build mode I had to add the portal sub route to the `package.json`
```json
{
  ...
  "homepage": "portal/.",
  ...
}
```
- In order to use react router dom, use the basename attribute when declaring the component

# TODOs:
- [x] Need to setup proxy to use dev servers
- [x] Ability to setup the dev servers from go if in debug
  -   Passing the dev server port and url in spa.Site
