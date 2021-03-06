
# Online White Board

# 起動

## サーバー
```
cd server
go run main.go
```

## クライアント
```
cd client
npm run dev
```

# 開発

## GoAPIサーバーソース自動生成
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
    
```

### 参考資料
- https://blog.masu-mi.me/post/2021/12/23/no-good-go-code-generator-of-openapi-v3.0/
- https://qiita.com/doriven/items/7422f565d6ad2e8ff956
- https://qiita.com/amuyikam/items/e8a45daae59c68be0fc8
- https://future-architect.github.io/articles/20210427c/


## TypeScriptクライアントソース自動生成
```
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
 
### 補足）テンプレート作成
- TypeScript3.8から ”Type-Only Imports and Export” が導入された
- しかしながら、openapi-generatorのデフォルトではこれに対応していない模様
- 対応として、「type」文字列を追記したテンプレートを「/openapi/template」配下に作成した
- 参考
  - https://www.typescriptlang.org/docs/handbook/release-notes/typescript-3-8.html
  - https://github.com/OpenAPITools/openapi-generator/issues/11179
  - https://qiita.com/watiko/items/0961287c02eac9211572
  - https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/typescript-axios


## Terraformでインフラ作成
- AWS上でリソースを作成後、Terraformのインポート機能を利用
  ```
  cd terraform
  terraform init
  terraform import -var-file="env.tfvars" aws_cloudfront_function.cloudfront_function arn:aws:cloudfront::749794220379:function/handler
  terraform plan -var-file="env.tfvars"
  terraform apply -var-file="env.tfvars"
  ```
  
- 備考
  - CloudFrontのFunctionをインポートしてもハングアップ状態になったため、リソースを削除しTerraformで新規作成をした



# デプロイ
## 自動デプロイ
- Githuh Actionsを利用

## 手動デプロイ
- クライアントソースデプロイ
  ```
  cd client
  aws s3 sync . s3://first-s3-jackyutajack --delete 
  ```

- サーバソースデプロイ
  ```
  cd server
  docker build -t first-ecr-jackyutajack .
  docker tag first-ecr-jackyutajack:latest 749794220379.dkr.ecr.ap-northeast-1.amazonaws.com/first-ecr-jackyutajack:latest
  docker push 749794220379.dkr.ecr.ap-northeast-1.amazonaws.com/first-ecr-jackyutajack:latest
  ```

