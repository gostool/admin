# Admin By GoFrame

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image`: Run `make build` to build image according `manifest/docker, and push image to your repository`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.

## [doc](https://uqgd478det.feishu.cn/docs/doccnT5x7PDb83huaxNQAjTtw9e#uitve9)
## [onlineDemo](https://api.pyhuo.top/admin/)
- userName: guest
- password: guest
### dist

```
npm run build
ln -s ~/github/web/dist dist
```