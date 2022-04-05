


swagger generate server -f ./swagger/swagger.yml -t server/gen

afdsaf
fsdf

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


https://blog.masu-mi.me/post/2021/12/23/no-good-go-code-generator-of-openapi-v3.0/

https://qiita.com/doriven/items/7422f565d6ad2e8ff956
https://qiita.com/amuyikam/items/e8a45daae59c68be0fc8





docker run --rm \
    -v $PWD:/local \
    -u "$(id -u $USER):$(id -g $USER)" \
    -v /etc/passwd:/etc/passwd:ro \
    -v /etc/group:/etc/group:ro \
    openapitools/openapi-generator-cli generate \
    -i /local/openapi/openapi.yaml \
    -g go-server \
    -o /local/server

