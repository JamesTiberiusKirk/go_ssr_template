#!/bin/sh

# Create Go APP
# BE CAREFUL USING THIS, IF U GIVE IT SOME WEIRD PATH IT MIGHT F*CK UP YOUR FILES
# !!!!!!!!! PATH SAFETY IS NOT IMPLEMENTED !!!!!!!!
# ./cga.sh -n test_project -p ~/Projects -as
# -n name of project
# -p path for the new project
# -a for adding API boilerplate
# -s for adding SSR boilerplate

PARENT_DIR=~/Projects

PROJECT_NAME=""
PROJECT_PATH=""
AUTH=""
SSR=false
API=false
STATIC=false

SITE_CONFIG_1="s := site.NewSite(config.Http.RootPath, db, sessionManager, e)"
SITE_CONFIG_2="  s.Serve()"

API_CONFIG_1="  a := api.NewApi(e.Group(config.Http.RootApiPath), config.Http.RootApiPath, db,"
API_CONFIG_2="  sessionManager)"
API_CONFIG_3="  a.Serve()"

MAIN_GO_SED_STR=''

ESCAPED_SITE_CONFIG=$(printf '%s\n' "$SITE_CONFIG_1" | sed -e 's/[\/&]/\\&/g')"\n"
ESCAPED_SITE_CONFIG=$ESCAPED_SITE_CONFIG$(printf '%s\n' "$SITE_CONFIG_2" | sed -e 's/[\/&]/\\&/g')"\n\n"

ESCAPED_API_CONFIG=$(printf '%s\n' "$API_CONFIG_1" | sed -e 's/[\/&]/\\&/g')"\n"
ESCAPED_API_CONFIG=$ESCAPED_API_CONFIG$(printf '%s\n' "$API_CONFIG_2" | sed -e 's/[\/&]/\\&/g')"\n"
ESCAPED_API_CONFIG=$ESCAPED_API_CONFIG$(printf '%s\n' "$API_CONFIG_3" | sed -e 's/[\/&]/\\&/g')"\n\n"

while getopts "hn:p:saSA:" arg; do
  case $arg in
    h)
      echo "Usage [-n:name] [-p:path] [-s] [-a] [-S]" 
      exit 0
      ;;
    n)
      PROJECT_NAME=$OPTARG
      PROJECT_PATH=$PARENT_DIR/$PROJECT_NAME
      ;;
    p)
      PARENT_DIR=$OPTARG
      ;;
    s)
      SSR=true
      ;;
    a)
      API=true
      ;;
    S)
      STATIC=true
      ;;
    A)
      AUTH=$OPTARG
      ;;
  esac
done

echo SSR $SSR
echo API $SSR
echo STATIC $STATIC
echo AUTH $AUTH

echo "creating new project"
echo $PROJECT_PATH

echo Template:  Go SSR Docker
SED_STR='s/go_ssr_template/'$PROJECT_NAME'/g'
mkdir $PROJECT_PATH

echo Copying base template files
cp ./init.go $PROJECT_PATH/
cp ./config.go $PROJECT_PATH/
cp ./README.md $PROJECT_PATH/
cp ./docker-compose.yml $PROJECT_PATH/
cp ./Dockerfile $PROJECT_PATH/
cp ./Makefile $PROJECT_PATH/
cp ./.gitignore $PROJECT_PATH/
cp ./.env $PROJECT_PATH/
cp -r ./session $PROJECT_PATH/
cp -r ./models $PROJECT_PATH/

if [[ $SSR == true ]]; then
  echo Copying SSR site folder
  cp -r ./site $PROJECT_PATH/
  MAIN_GO_SED_STR="$MAIN_GO_SED_STR$ESCAPED_SITE_CONFIG"
fi

if [[ $API == true ]]; then
  echo Copying API folder
  cp -r ./api $PROJECT_PATH/
  MAIN_GO_SED_STR="$MAIN_GO_SED_STR$ESCAPED_API_CONFIG"
fi

if [[ $STATIC == true ]]; then
  echo No support for static sites yet
fi

# if [[ $AUTH == "session" ]]; then
# fi

cp ./main.go.template $PROJECT_PATH/main.go
sed -i '' "s/{{code}}/$MAIN_GO_SED_STR/g" $PROJECT_PATH/main.go


cd $PROJECT_PATH

find . -type f -exec sed -i '' $SED_STR {} +
git init
go mod init $PROJECT_NAME
go mod tidy 
go get ./...

# echo Run go mod init and go get ./...

