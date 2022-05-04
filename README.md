
# Online White Board

## 起動

### サーバー
```
cd server
go run main.go
```

### クライアント
```
cd client
npm run dev
```

## 開発
```
swagger generate server -f ./swagger/swagger.yml -t server/gen

docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/petstore.yaml \
    -g go \
    -o /local/out/go

docker run --rm \
    -v $PWD:/local openapitools/openapi-generator-cli generate \
    -i /local/swagger/swagger.yml \
    -g go \
    -o /local/server/gen
```

- https://blog.masu-mi.me/post/2021/12/23/no-good-go-code-generator-of-openapi-v3.0/
- https://qiita.com/doriven/items/7422f565d6ad2e8ff956
- https://qiita.com/amuyikam/items/e8a45daae59c68be0fc8
- https://future-architect.github.io/articles/20210427c/


```
docker run --rm \
    -v $PWD:/local \
    -u "$(id -u $USER):$(id -g $USER)" \
    -v /etc/passwd:/etc/passwd:ro \
    -v /etc/group:/etc/group:ro \
    openapitools/openapi-generator-cli generate \
    -i /local/openapi/openapi.yaml \
    -g go-server \
    -o /local/server \
    --additional-properties=featureCORS=true
    

docker run --rm \
  -v $PWD:/app \
  -u "$(id -u $USER):$(id -g $USER)" \
  -v /etc/passwd:/etc/passwd:ro \
  -v /etc/group:/etc/group:ro \
  openapitools/openapi-generator-cli generate \
  -i /app/openapi/openapi.yaml \
  -g typescript-axios \
  -t /app/openapi/template \
  -o /app/client/src/gen
 ```
 
## クライアントジェネレート
### テンプレート作成
- TypeScript3.8から ”Type-Only Imports and Export” が導入された
- しかしながら、openapi-generatorのデフォルトではこれに対応していない模様
- type を加えたテンプレートを作成した

- https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-8.html
- https://github.com/OpenAPITools/openapi-generator/issues/11179
- https://qiita.com/watiko/items/0961287c02eac9211572
- https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/typescript-axios



```
cd server
go run ./main

docker build --network=host -t openapi .
docker run --rm -it openapi

```

## Terraform

```
cd terraform
terraform init
terraform import -var-file="env.tfvars" aws_cloudfront_function.cloudfront_function arn:aws:cloudfront::749794220379:function/handler

terraform plan -var-file="env.tfvars"
terraform apply -var-file="env.tfvars"

```

- CloudFrontのFunctionをインポートしてもハングアップ状態になった



## Deploy

```
aws s3 sync . s3://projectname-static-files --delete --grants read=uri=http://acs.amazonaws.com/groups/global/AllUsers --exclude "README.md"
aws s3 sync . s3://first-s3-jackyutajack --delete 

```
