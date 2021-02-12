# 【TIPS#7】 TermuxでWEBアプリを作ってアクセスしてみる（Go）

## はじめに

前回はアンドロイドスマホのTermuxアプリのLinux環境で、Python言語で書いたシンプルなWEBアプリケーションを起動し、アンドロイドスマホやPCのブラウザからアクセスし、Zaifから取得したXEMの最終価格を表示しました。

Termuxは、Android端末エミュレータおよびLinux環境アプリです。

今回は同様のことをGo言語で書いたシンプルなWEBアプリケーションで行ってみます。

Go言語でも同じことをするのは、Go言語がPython同様に非常に人気のあるプログラミング言語だからです。
私が個人的に一番好きな言語だからという理由もありますｗ。

TermuxアプリのLinux環境にGo言語の開発環境を導入する手順は、過去記事で紹介していますので、そちらを参照してください。

## WEBアプリケーションの起動準備

任意の場所にディレクトリ（フォルダ）を作成し、そこに移動します。

mkdir gowebapp
cd gowebapp

go mod init golang

`golang`の部分は好きな名前に変更してもよいです。

## WEBアプリの起動

go run app.go

またはビルドしてから実行する場合は、

go build app.go
./app

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

