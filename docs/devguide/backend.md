# バックエンド

## 場所

`be` ディレクトリにバックエンド関連のコードを入れています。

## 実行

開発用にサーバを起動するにはターミナルで以下のコマンドを実行します:

```shell
$ npm run start
```

その後 `http://localhost:8000/` にアクセスしてください (別の URL が表示されますが、無視してください)。
ただし、8000 番ポートが使用できなかった場合、1333 番ポートにフォールバックされるので、
8000番ポートでアクセスできなかった場合はこちらを試してください。

サーバを止めるにはターミナルの画面で Control+C を押します。

## テスト

無理しなくてもよいですが、ユニットテストを書いていくと毎回ブラウザで確認する手間が省けて開発が捗るかもしれません。
書き方は既存の `*_test.go` のファイルを参考にしてください。

テストを実行したい場合、ターミナルで

```shell
$ go test ./...
```

を実行するとテストが動きます。

## コードのフォーマット

コミットする前に以下のコマンドを実行し、コードの体裁を整えます。

```shell
$ go fmt ./...
```
