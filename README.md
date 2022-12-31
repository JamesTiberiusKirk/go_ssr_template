# Create Go APP
## BE CAREFUL USING THIS, IF U GIVE IT SOME WEIRD PATH IT MIGHT F*CK UP YOUR FILES
- !!!!!!!!! PATH SAFETY IS NOT IMPLEMENTED !!!!!!!!

## Usage
`./cga.sh -n test_project -p ~/Projects -as`  
- `-n` name of project  
- `-p` path for the new project  
- `-a` for adding API boilerplate  
- `-s` for adding SSR boilerplate

## Featuring
- Mostly out of your way extensible framework 
- Echo framework 
- Goview to help with templating
- GORM with Postgres
- Session authentication

## TODOs:
- [ ] Link all of the libraries in "featuring" section
- [ ] Actually setup the rootSitePath
- [ ] Setup a good way to serve static content such as libraries
- [ ] Maybe think about setting up some theme or CSS library?
- [ ] Maybe find an actual way to load up the .env file in the dev_run.sh
- [x] Make some CLI or script for easy deployment of this Template
  - Not much needed other then copying the folder then using `sed -i` to replace the name to the desired one
- [ ] Look more at making a dev container
  - Maybe there's some easy way to make it portable across x86 and arm
- [x] Standalone API handler framework (like site but without the renderer)
- [x] Rethink the GetXMethod handler system
  - It adds too much boilerplate 
- [ ] Maybe think of a way to hardcode in the path to be used outside each page for redirection
  - Probably just consts 
- [x] Add a menu system to frame and pages
  - [ ] Just need to actually create a menu in the templates 
- [ ] Add support for SPAs (vue, angular, react, etc...) inside pages
- [ ] Add support for WASM inside pages
  - Maybe look at one of the WASM frameworks out there for go
- [] Re-write the cga script so it actually works properly
  - Seems to be very jank with it being just a shell script, maybe re-write it in go? 
  - Bugs probs still exist
  - No safety exists at the moment
    - Need to quit the script the moment there seems to be something wrong to avoid f*ing problems


