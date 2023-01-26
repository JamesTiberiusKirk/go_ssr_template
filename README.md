# Create Go APP

## Featuring
- Mostly out of your way extensible framework 
- Echo framework 
- (Goview)[https://github.com/foolin/goview] to help with templating
  - Forked a basic version of it in this repo
- GORM with Postgres
- Session authentication
- (Reflex)[github.com/cespare/reflex] for hotloading
- Forked (tygo)[github.com/gzuidhof/tygo] and added a few features for Typescript type generation from go structs

# ROADMAP
## V0.1.0
- Version 0.1 will be about having all the bells and whistles that anyone might ever want
- If the user does not want some of those features, they can just remove them manually 

### TODOs:
- [x] Link all of the libraries in "featuring" section
- [x] Actually setup the rootSitePath
- [x] Setup a good way to serve static content such as libraries
- [x] Setup a 404 default re-routing
- [ ] Add support for SPAs (vue, angular, react, etc...) inside pages
  - With support for letting the spa use its own router
- [ ] Finish a basic version of the cli
- [ ] Package up the cli
- [ ] Figure out how to remove the exports from js dist
- [ ] Implement some form of a TUI around this to pick options and selections (Bubble Tea)

## V0.2.0
- This will be about expanding on the feature set

### TODOs:
- [ ] Add support for web sockets
- [ ] Abstract the db behind a struct so the client can more easily swapped out
- [ ] Add an internal server failure page
- [ ] Add support for WASM inside pages
  - Maybe look at one of the WASM frameworks out there for go

## Template TODOs:
- [x] Link all of the libraries in "featuring" section
- [ ] Actually setup the rootSitePath
- [ ] Setup a good way to serve static content such as libraries
- [ ] Maybe think about setting up some theme or CSS library?
- [ ] Maybe find an actual way to load up the .env file in the dev_run.sh
- [x] Make some CLI or script for easy deployment of this Template
  - Not much needed other then copying the folder then using `sed -i` to replace the name to the desired one
- [x] Standalone API handler framework (like site but without the renderer)
- [x] Rethink the GetXMethod handler system
  - It adds too much boilerplate 
- [x] Maybe think of a way to hardcode in the path to be used outside each page for redirection
  - Probably just consts 
- [x] Add a menu system to frame and pages
- [ ] Add support for SPAs (vue, angular, react, etc...) inside pages
- [ ] Add support for WASM inside pages
  - Maybe look at one of the WASM frameworks out there for go
- [ ] Add support for web sockets
- [x] EITHER replace echoview or re-write it myself, it's been a year since last commit and its missing some needed features
  - Probably the best thing would be to implement it myself within my own site struct with complete support for stuff like:
    - [x] default 404s 
    - [x] options to exclude master frame per page 
    - [x] custom template file includes
    - [x] render function for the frame (so that base data can be included such as authed user, altho might still be better done in the GetPageHandler function instead)
- [x] Add typescript support to the js proto framework
  - [x] Go struct transpiration to typescript
  - [x] Typescript typings for the proto framework 
  - [ ] Try to implement the base framework functions in TS
  - [ ] Get rid of lint/lsp warning for unused declarations in TS
- [ ] Possibly throw the page status code in the get page data function so that each page can return their own status codes
- [ ] Add an internal server failure page

# For getting the dev script to run 
```sh 
go install github.com/cespare/reflex@latest
go install github.com/JamesTiberiusKirk/tygo@v0.2.5
```
