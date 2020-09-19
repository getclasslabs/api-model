# api-model
{DESCRIPTION OF THE API}

## Routes
{API ROUTES WITH BODY AND RESPONSE EXAMPLES}

## How to test
{makefile instructions}

## Domain
{Description of the domain the api controls}

## Model to api
To turn this model into an API you'll have to: 
1) Edit all the references of: `{apiname}` to the name of your api.
2) Edit the config files (docker-config.yaml and config.yaml) with the things that your api will need. (optional)
3) Change the name of the dir: ./cmd/{apiname} with the name of your api.
4) Execute `go mod init github.com/getclasslabs/{apiname}`
5) Put everything on vendor with: `go mod vendor`
6) You're ready to go!
