# workspace admin SDK

## what is this sample

- use service account to call workspace admin API with domain-wide delegation
- JWT client with service account key file

## warning

- do not forget change the `jwt.Config.Subject` with your deligated workspace domain email

## ref

- https://support.google.com/a/answer/162106
- https://developers.google.com/identity/protocols/oauth2/service-account#delegatingauthority
- https://developers.google.com/identity/protocols/oauth2/service-account
- https://developers.google.com/admin-sdk/reports/v1/quickstart/go
- https://developers.google.com/admin-sdk/directory/v1/guides/delegation
- https://developers.google.com/admin-sdk/reports/v1/guides/delegation
- https://developers.google.com/admin-sdk/reports/v1/guides/authorizing
- https://developers.google.com/workspace/guides/create-credentials
