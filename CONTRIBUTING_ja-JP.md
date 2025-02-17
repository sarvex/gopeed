# Gopeed コントリビューターガイド

まず最初に、Gopeed への貢献に興味を持っていただきありがとうございます。このガイドは、あなたが Gopeed の
開発に参加するための手助けとなるでしょう。

## ブランチの説明

このプロジェクトのメインブランチは `main` ブランチのみです。Gopeed の開発に参加したい場合は、
まずこのプロジェクトをフォークし、フォークしたプロジェクトで開発を行ってください。開発が完了したら、
このプロジェクトに PR を提出し、`main` ブランチにマージしてください。

## ローカル開発

開発およびデバッグはウェブ上で行うことを推奨する。まずバックエンドのサービスを起動し、
コマンドライン `go run cmd/api/main.go` で起動する。サービスのデフォルトポートは `9999` で、
次にフロントエンドの flutter プロジェクトを `debug` モードで起動して実行します。

## 翻訳

Gopeed の国際化ファイルは `ui/flutter/assets/locales` ディレクトリにあります。
このディレクトリに対応する言語ファイルを追加するだけでよいです。


ロケール編集後にロケールを生成:


```
get generate locales
```

翻訳については `en_us.dart` を参照してください。

## flutter での開発

コミットする前に `dart format ./ui/flutter` を実行し、コードを標準の dart フォーマットにしておくことを忘れないでください

api/models を編集したい場合は build_runner watcher をオンにします:

```
flutter pub run build_runner watch
```


get-cli コマンドの使用法:

```
 create:
    controller:  Generate controller
    page:  Use to generate pages
    view:  Generate view
  generate:
    locales:  Generate translation file from json files
    model:  generate Class model from json
```
