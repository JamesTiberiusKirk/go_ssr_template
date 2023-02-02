# Create Go APP

We are at v0.1.0 of this template.
If curious you can see my [roadmap here](./ROADMAP.md)

## Featuring
- Mostly out of your way extensible framework 
- Echo framework 
- [Goview](https://github.com/foolin/goview) to help with templating
  - Forked a basic version of it in this repo
- GORM with Postgres
- Session authentication
- [Reflex](github.com/cespare/reflex) for hotloading
- Forked [tygo](github.com/gzuidhof/tygo) and added a few features for Typescript type generation from go structs
  - [Fork here](github.com/JamesTiberiusKirk/tygo)

## Template Architecture Features:
- [API](api/)
- [Site](site/)
  - [SPA hosting](site/spa/)
  - [SSR/Templating](site/page/)
    - [JS/TS templating framework with full stack type safety](site/page#jsts-minimal-ssr-framework)

# For getting the dev script to run 
```sh 
go install github.com/cespare/reflex@latest
go install github.com/JamesTiberiusKirk/tygo@v0.2.5
```
