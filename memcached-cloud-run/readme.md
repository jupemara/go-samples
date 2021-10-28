# memcached on cloud run with vpc connector

## deployment

### authentication

```sh
# set user credential as application default credential
$ gcloud auth application-default login
# set artifact registry credential
$ gcloud auth configure-docker ${YOUR_ARTIFACT_REGISTROY_PASS_MINUS_PREFIX}
```

#### artifact registory pass

like `${REGION}-docker.pkg.dev/${GCP_PROJECT_ID}/${YOUR_SPECIFIED_PREFIX}`.
e.g: `asia-northeast1-docker.pkg.dev/jupemara/d`

### with ko

```sh
$ export KO_DOCKER_REPO=${YOUR_ARTIFACT_REGISTORY_PASS}
$ gcloud run deploy ${YOUR_APPLICATION_NAME} \
  --image=$(ko publish .) \
  --platform=managed \
  --region=${REGION} \
  --vpc-connector=${YOUR_VPC_CONNECTOR_NAME}
```
