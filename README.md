# Create Go APP
## BE CAREFUL USING THIS, IF U GIVE IT SOME WEIRD PATH IT MIGHT F*CK UP YOUR FILES
- !!!!!!!!! LOCAL PATH SAFETY IS NOT IMPLEMENTED !!!!!!!!

## Featuring
- Mostly out of your way extensible framework 
- Echo framework 
- (Goview)[https://github.com/foolin/goview] to help with templating
  - Although now I've had to pretty much fork goview and modify it to suit my needs
- GORM with Postgres
- Session authentication

## Template TODOs:
- [ ] Link all of the libraries in "featuring" section
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
  - [ ] Just need to actually create a menu in the templates 
- [ ] Add support for SPAs (vue, angular, react, etc...) inside pages
- [ ] Add support for WASM inside pages
  - Maybe look at one of the WASM frameworks out there for go
- [x] Re-write the cga script so it actually works properly
  - Seems to be very jank with it being just a shell script, maybe re-write it in go? 
  - Bugs probs still exist
  - No safety exists at the moment
    - Need to quit the script the moment there seems to be something wrong to avoid f*ing problems
- [ ] Look at echoview partials for including files such as the SSR library and potentially even `.js/.css/.any` files in normal pages
- [ ] EITHER replace echoview or re-write it myself, it's been a year since last commit and its missing some needed features
  - Probably the best thing would be to implement it myself within my own site struct with complete support for stuff like:
    - [ ] default 404s 
    - [ ] options to exclude master frame per page 
    - [x] custom template file includes
    - [ ] render function for the frame (so that base data can be included such as authed user, altho might still be better done in the GetPageHandler function instead)

## CRA CLI TODOs:
- [ ] Cleanup the vars section
- [ ] Organise the functions to be bound to an instance of either the options struct or another struct which would hold the options
  - [ ] Have a lot of the hard-coded global variables inside that function to make it more modular and to have a lot of that config organised in one place
  - [ ] Donno if it makes sense but could even have that config being pulled from an env file...
- [ ] Implement some form of a TUI around this to pick options and selections (Bubble Tea)
