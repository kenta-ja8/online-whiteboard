


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
