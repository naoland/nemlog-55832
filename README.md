# 【TIPS#7】 TermuxでWEBアプリを作ってアクセスしてみる（Go）

## はじめに

前回はアンドロイドスマホのTermuxアプリのLinux環境で、Python言語で書いたシンプルなWEBアプリケーションを起動し、アンドロイドスマホやPCのブラウザからアクセスし、Zaifから取得したXEMの最終価格を表示しました。

`Termux`というのは、Android端末エミュレーターおよびLinux環境アプリです。

今回は前回の記事同様のことを、Go言語で実装したシンプルなWEBアプリケーションについて紹介します。

Go言語でも同じことをするのは、Go言語がPython同様、非常に人気のあるプログラミング言語だからです。

私が個人的に一番好きな言語だからという理由もありますｗ。

TermuxアプリのLinux環境にGo言語の開発環境を導入する手順は、過去記事で紹介していますので、そちらを参照してください。

## WEBアプリケーションの起動準備

任意の場所にディレクトリ（フォルダー）を作成し、そこに移動します。

次のコマンドをタイプして実行してください。

`gowebapp`というディレクトリ（フォルダー）作成します。

```
$ mkdir gowebapp
```

`gowebapp`というディレクトリ（フォルダー）に移動します。


```
$ cd gowebapp
```

Goのモジュールやパッケージの依存関係を管理するために初期化を行います。

```
$ go mod init naoland
```

実行例

```
main $ go mod init naoland
go: creating new go.mod: module naoland

```

今回のような簡単なケースでは、`naoland`の部分は好きな名前に変更してもかまいません。


WEBアプリを作るために必要なパッケージをインストールします。

```
$ go get -u github.com/gin-gonic/gin
```

実行例

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
Ginが依存しているパッケージがインストールされます。


## WEBアプリの起動

```
$ go run app.go
```

またはビルドしてから実行する場合は、

```
$ go build app.go
./app
```

次にビルドされたアプリを起動します。

```
$ ./app
```

上記の手順を実行すると、`gowebapp`というWEBアプリが起動します。


## WEBアプリの動作確認

アンドロイドスマホ上のChromeアプリで確認する方法と、PC上のChromeアプリで確認する方法があります。

### アンドロイドスマホ上のChromeアプリで確認する方法

- Chromeブラウザアプリを起動します。
- http://localhost:8080/last_price/xem_jpy にアクセスします。

次の結果が得られるはずです。




### PC上のChromeアプリで確認する方法


- Chromeブラウザアプリを起動します。
- http://localhost:8080/last_price/xem_jpy にアクセスします。

次の結果が得られるはずです。

## ソースコード
```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseZaifInfo Zaifからのレスポンス
type ResponseZaifInfo struct {
	LastPrice float64 `json:"last_price"`
}

func main() {
	r := gin.Default()
	r.GET("/last_price", func(c *gin.Context) {
		price, err := fetchLastPrice()
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, "XEM最終価格: %.3f", price)
	})
	r.Run("0.0.0.0:3000")
}

// Zaif取引所の最終価格を取得して、価格を返します
func fetchLastPrice() (price float64, err error) {
	uri := "https://api.zaif.jp/api/1/last_price/xem_jpy"
	req, _ := http.NewRequest("GET", uri, nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	response := new(ResponseZaifInfo)
	err = json.Unmarshal(byteArray, &response)

	if err != nil {
		log.Printf("マーシャルエラー: %#v\n", err) // 構造調べ
		return
	}
	price = response.LastPrice
	return
}
```

## まとめ

いかがでしたでしょうか？

前回の記事でご紹介したPython言語によるコードを実行した場合よりも早く結果が得られたのではないでしょうか？
Go言語で作成したプログラムは高速だと言われており、速度が重要なサービスの開発で使用されるケースが増えています。
Go言語の文法などは非常にシンプルなので、慣れればPythonよりもコードを書くのが楽だと思います。
Python言語は簡単だと言われますが、それはバージョンが2.x時代であって、現在は非常に複雑な言語の1つだと思います。
また1つの事を実現するコードは大体同じになると言われてますが、それも過去の事です。現在は同じことを実現するコードをいろんな方法で書くことができてしまいます。

その点、Go言語は昔のPythonのように、1つのことを実現する方法のコードは大体同じになります。

国家試験でPythonを選択できるようになりましたが、馬鹿じゃないのか？と思います。時代が変わったといえ、スクリプト言語なんかを国家試験レベルの対象にすべきではないと私は思ってます。

ということで、今後のTipsシリーズではPython言語も扱いますが、わかりやすい場合はGo言語で紹介したいと思います。


## 関連情報へのリンク

- [techbureau/zaifapi: zaifのAPIを簡単にコール出来るようにしました。](https://github.com/techbureau/zaifapi)
- [gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.](https://github.com/)
- [ZaifAPI ドキュメント — Zaif api document v1.1.1 ドキュメント](https://techbureau-api-document.readthedocs.io/ja/latest/index.html)
- https://scrapbox.io/api/code/naoland/mypage/app.go

