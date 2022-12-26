# Template for quickly setting up go SSR site with echo and goview

## Featuring
- Mostly out of your way extensible framework 
- Echo framework 
- Goview to help with templating
- GORM with Postgres
- Session authentication

## Usage
- Copy to a directory of your choosing
- Run the following command just replacing `YOURNAMEHERE` with your desired name
```sh
find . -type f -not -path './db-mount/*' -not -path './.git/*' -exec sed -i '' 's/go_ssr_template/YOURNAMEHERE/g' {} \;
```
- If its a clean repo just the simpler sed only command will work also
```sh 
sed -i '' 's/music_manager/go_ssr_template/g' ./* 
```

## TODOs:
- [] Link all of the libraries in "featuring" section
- [] Actually setup the rootSitePath
- [] Setup a good way to serve static content such as libraries
- [] Maybe think about setting up some theme or CSS library?
- [] Maybe find an actual way to load up the .env file in the dev_run.sh
- [] Make some CLI or script for easy deployment of this Template
  - Not much needed other then copying the folder then using `sed -i` to replace the name to the desired one
- [] Look more at making a dev container
  - Maybe there's some easy way to make it portable across x86 and arm
- [] Standalone API handler framework (like site but without the renderer)
- [x] Rethink the GetXMethod handler system
  - It adds too much boilerplate 
- [] Maybe think of a way to hardcode in the path to be used outside each page for redirection
  - Probably just consts 
- [] Add a menu system to frame



