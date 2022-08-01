# Images API on AppEngine 2nd gen

## build and deploy

```sh
$ gcloud app deploy
```

then you can get the serving_url on appengine gen2:tada::tada::tada:

## how to setup?

1. setup fine-grained Access Control Lists "きめ細かいアクセス制御リスト" not uniform bucket-level
2. add `app_engine_apis: true` to `app.yaml`

## links

- https://cloud.google.com/appengine/docs/standard/go111/images
- https://cloud.google.com/appengine/docs/standard/go/go-differences#writing_a_main_package
