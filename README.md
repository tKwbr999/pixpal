# PixPal

PixPalはピクセルキャラクターを生成するGo製のツールです。

![Sample SVG](sample.svg)



## 利用方法

`example`ディレクトリに生成例があります。

```bash
cd example/defaultMan
go run main.go
```

これにより、`example/defaultMan`ディレクトリに`default_man.svg`が生成されます。


## コマンドラインツールの使い方

### インストール

```bash
go install github.com/tKwbr999/pixpal@latest
```

### 実行

```bash
pixpal
```
これにより、`default_man.svg`という名前のSVGファイルが生成されます。

または、特定のファイル名でSVGファイルを生成する場合
```bash
pixpal -o custom_name.svg
```
