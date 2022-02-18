# Workload Identity + Cloud SQL Auth Proxy + IAM Authentication

Workload Identity を有効化した GKE Cluster に対して、 Cloud SQL Auth Proxy をサイドカーで動かして、IAM 認証を行うためのコード。

## warning

2022-02-18 の段階では、 Cloud SQL Auth Proxy の自動認証は MySQL に対応しておらず、 PostgreSQL のみとなる。ので、一旦 IAM 認証なしで、 Cloud SQL Auth Proxy 専用ユーザで動かすことにしますねｗ
Workload Identity を有効にした GKE の metadata から token を取得して動かすこともできなくはないけど、 token の expiration が 1 時間なので、都度自分で取得するのはめんどいかもしれない

## サンプルの流れ

1. Workload Identity を有効にした GKE Cluster を作る
2. CloudSQL の instance 雑に作る
3. CloudSQL のユーザに Cloud SQL Proxy Auth Proxy 専用ユーザを作る

```sh
mysql > create user 'cloudsql-proxy-special'@'cloudsqlproxy~%';
mysql > grant show databases on *.* to 'cloudsql-proxy-special'@'cloudsqlproxy~%';
```

このとき、そもそもの CloudSQL instance への接続を Cloud SQL Auth Proxy 経由で行いたい場合は

```sh
# ローカルの docker 上で Cloud SQL Auth Proxy を立ち上げる
$ docker run -it -p 3306:3306 \
  -v gcp-service-account-key.json:/config \
  gcr.io/cloudsql-docker/gce-proxy /cloud_sql_proxy \
  -instances=GCP_PROJECT_ID:GCP_REGION:GCP_CLOUDSQL_INSTANCE_NAME=tcp:0.0.0.0:3306 -credential_file=/config
$ mysql -u root -p -h 127.0.0.1
```

ローカルの docker 上に立ち上げた Cloud SQL Auth Proxy 経由で mysql にアクセスする。

4. そろそろ GKE Cluster が立ち上がってる頃合いなので、 k8s 上の SA と GCP 上の SA を binding する

```sh
# GKE の SA と GCP の SA を bindig
$ gcloud iam service-accounts add-iam-policy-binding \
  cloudsql-proxy@GCP_PROJECT_ID.iam.gserviceaccount.com \
  --role roles/iam.workloadIdentityUser \
  --member="serviceAccount:GCP_PROJECT_ID.svc.id.goog[K8S_NAME_SPACE/K8S_SERVICE_ACCOUNT]"
```

5. app.yaml を deploy する

## To-Do

- [ ] MySQL の IAM 自動認証がきたら再度検証
- [ ] metadata から token 取得して 1 時間で refresh する (latency にシビアじゃないアプリならコレもあり)
