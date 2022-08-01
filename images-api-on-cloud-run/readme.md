# Images API が Cloud Run で試験するコード

## docker auth with gcloud

```sh
$ gcloud auth configure-docker
```

## deploy

just do that!!

```sh
# this is sample docker repository name
$ export KO_DOCKER_REPO=asia-northeast1-docker.pkg.dev/awesome-gcp/docker
$ gcloud run deploy images-api-on-cloud-run \
  --image=$(ko publish .) \
  --allow-unauthenticated \
  --region=asia-northeast1
```

## links

- https://cloud.google.com/appengine/docs/standard/runtimes
- https://cloud.google.com/appengine/docs/standard/go/blobstore/reference#BlobKeyForFile

## result

`not an App Engine context` error message will be appeared on your screen...
