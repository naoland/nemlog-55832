# 【TIPS#7】 TermuxでWEBアプリを作ってアクセスしてみる（Go）

## はじめに

前回はアンドロイドスマホのTermuxアプリのLinux環境で、Python言語で書いたシンプルなWEBアプリケーションを起動し、アンドロイドスマホやPCのブラウザからアクセスし、Zaifから取得したXEMの最終価格を表示しました。

Termuxは、Android端末エミュレーターおよびLinux環境アプリです。

今回は同様のことをGo言語で書いたシンプルなWEBアプリケーションで行ってみます。

Go言語でも同じことをするのは、Go言語がPython同様、非常に人気のあるプログラミング言語だからです。
私が個人的に一番好きな言語だからという理由もありますｗ。

TermuxアプリのLinux環境にGo言語の開発環境を導入する手順は、過去記事で紹介していますので、そちらを参照してください。

## WEBアプリケーションの起動準備

任意の場所にディレクトリ（フォルダー）を作成し、そこに移動します。

```
mkdir gowebapp
cd gowebapp
```
```
go mod init golang
05:39:45 (.myvenv) nao@330 nemlog-55832 ±|main ✗|→ go mod init naoland
go: creating new go.mod: module naoland

```
`golang`の部分は好きな名前に変更してもかまいません。


パッケージをインストールします。

```
main $ go get -u github.com/gin-gonic/gin
go: downloading github.com/gin-gonic/gin v1.6.3
go: github.com/gin-gonic/gin upgrade => v1.6.3
go: downloading github.com/gin-contrib/sse v0.1.0
go: downloading github.com/go-playground/validator/v10 v10.2.0
go: downloading github.com/ugorji/go v1.1.7
go: downloading github.com/golang/protobuf v1.3.3
go: downloading github.com/json-iterator/go v1.1.9
go: downloading golang.org/x/sys v0.0.0-20200116001909-b77594299b42
go: downloading github.com/go-playground/universal-translator v0.17.0
go: downloading github.com/leodido/go-urn v1.2.0
go: downloading github.com/ugorji/go/codec v1.1.7
go: downloading github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421
go: downloading github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742
go: downloading github.com/go-playground/locales v0.13.0
go: github.com/modern-go/reflect2 upgrade => v1.0.1
go: github.com/leodido/go-urn upgrade => v1.2.1
go: github.com/json-iterator/go upgrade => v1.1.10
go: gopkg.in/yaml.v2 upgrade => v2.4.0
go: github.com/modern-go/concurrent upgrade => v0.0.0-20180306012644-bacd9c7ef1dd
go: github.com/golang/protobuf upgrade => v1.4.3
go: golang.org/x/sys upgrade => v0.0.0-20210124154548-22da62e12c0c
go: github.com/ugorji/go/codec upgrade => v1.2.4
go: github.com/go-playground/validator/v10 upgrade => v10.4.1
go: downloading github.com/json-iterator/go v1.1.10
go: downloading golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c
go: downloading github.com/go-playground/validator/v10 v10.4.1
go: downloading gopkg.in/yaml.v2 v2.4.0
go: downloading github.com/ugorji/go v1.2.4
go: downloading github.com/golang/protobuf v1.4.3
go: downloading github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
go: downloading github.com/modern-go/reflect2 v1.0.1
go: downloading github.com/leodido/go-urn v1.2.1
go: downloading github.com/ugorji/go/codec v1.2.4
go: downloading google.golang.org/protobuf v1.23.0
go: google.golang.org/protobuf upgrade => v1.25.0
go: golang.org/x/crypto upgrade => v0.0.0-20201221181555-eec23a3978ad
go: downloading google.golang.org/protobuf v1.25.0
go: downloading golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
```
Ginは大きめのWEBフレームワークなので、けっこうたくさんのパッケージがインストールされます。


## WEBアプリの起動

```
go run app.go
```

またはビルドしてから実行する場合は、

```
go build app.go
./app
```

## WEBアプリの動作確認

## ソースコード
```go
  package main

  import (
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      uri := "https://api.zaif.jp/api/1/last_price/btc_jpy"
      req, _ := http.NewRequest("GET", uri, nil)

      client := new(http.Client)
      resp, _ := client.Do(req)
      defer resp.Body.Close()

      byteArray, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(byteArray))
  }
  // 結果 {"last_price": 130065.0}
```

## まとめ

## 関連情報へのリンク

- [techbureau/zaifapi: zaifのAPIを簡単にコール出来るようにしました。](https://github.com/techbureau/zaifapi)
- [gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.](https://github.com/)
- [ZaifAPI ドキュメント — Zaif api document v1.1.1 ドキュメント](https://techbureau-api-document.readthedocs.io/ja/latest/index.html)
- https://scrapbox.io/api/code/naoland/mypage/app.go

