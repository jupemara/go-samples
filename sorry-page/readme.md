# sorry ページを IP Address で特定のレスポンス

## overview

- Plan A
  1. header base routing で特定の backend に流す
  1. もし可能なら X-Forwarded-For を指定できれば完璧: これはできませんでした。不可能
- Plan B
  1. cloud armor で IP アドレスで指定した URL に redirect

## 手順: Plan A

1. コンテンツ用の LB を作成: 受け取ったヘッダをすべて垂れ流すコンテナを cloud run で backend にする
1. sorry 用の LB を作成: 1 で垂れ流されたヘッダーでルーティングする
   1. X-Forwarded-For でできれば FIN
   1. できなければ適当なヘッダーでルーティングする
