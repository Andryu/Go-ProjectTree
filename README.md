# Go-ProjectTree

## Go言語のAPIプロジェクトツリーを考えてみた(フレームワークなし)

### ディレクトリ構成

  - main.go
  - config
    - config.go

  - models
    - user.go

  - services
    - ranking.go

  - controllers
    - v1_controller.go

### 説明

- config  ......環境変数などの設定周りを初期化

- main.go ......プロジェクトのメイン関数を定義

- controllers ...コントローラー, ロジック結果内容をログ出力

- models  ......モデル

- services ......Viewモデル+ビジネスロジック

