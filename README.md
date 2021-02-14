# 【TIPS#7】 TermuxでWEBアプリを作ってアクセスしてみる（Go）

## はじめに

前回はアンドロイドスマホのTermuxアプリのLinux環境で、Python言語で書いたシンプルなWEBアプリケーションを起動し、アンドロイドスマホやPCのブラウザからアクセスし、Zaifから取得したXEMの最終価格を表示しました。

`Termux`というのは、Android端末エミュレーターおよびLinux環境アプリです。

今回は前回の記事同様のことを、Go言語で実装したシンプルなWEBアプリケーションについて紹介します。

Go言語でも同じことをするのは、Go言語がPython同様、非常に人気のあるプログラミング言語だからです。

私が個人的に一番好きな言語だからという理由もありますｗ。

TermuxアプリのLinux環境にGo言語の開発環境を導入する手順は、過去記事で紹介していますので、そちらを参照してください。

## WEBアプリケーションの起動準備

今回の記事用にGitHubにコードを用意していますので、それを使います。

https://github.com/naoland/nemlog-55832

任意の場所にディレクトリ（フォルダー）で次のコマンドを実行して、ソースコードなどが含まれたリポジトリをダウンロードしてください。

```
$ git clone https://github.com/naoland/nemlog-55832.git
```

もしgitがインストールされておらずエラーが表示される場合は、`git`をインストールします。
インストール方法は過去記事でも紹介していますが、簡単なのでここでも一応ご紹介しておきます。

次のコマンドをタイプして実行してください。

```
$ pkg install git
```

前述のコマンドを実行すると`nemlog-55832`というフォルダーができているはずなので、移動します。

次のコマンドをタイプして実行してください。

```
$ cd nemlog-55832
```

WEBアプリをビルドします。


```
$ make build
```

実行例

```
$ make build
go: downloading github.com/gin-gonic/gin v1.6.3
go: downloading github.com/mattn/go-isatty v0.0.12
go: downloading github.com/gin-contrib/sse v0.1.0
go: downloading github.com/go-playground/validator/v10 v10.4.1
go: downloading github.com/golang/protobuf v1.4.3
go: downloading gopkg.in/yaml.v2 v2.4.0
go: downloading github.com/ugorji/go v1.2.4
go: downloading golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c
go: downloading github.com/ugorji/go/codec v1.2.4
go: downloading golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
go: downloading github.com/leodido/go-urn v1.2.1
go: downloading github.com/go-playground/universal-translator v0.17.0
go: downloading google.golang.org/protobuf v1.25.0
go: downloading github.com/go-playground/locales v0.13.0
```

Ginが依存しているパッケージが自動でインストールされた後にビルドされます。

もし`make`がインストールされていない場合は、次のようにタイプしてインストールします。

```
$ pkg install make
```

どのようにアプリをビルドするかなど、`Makefile`の内容を参照してください。

もしWindows用にビルドしたい場合は次のコマンドを実行してください。

```
$ make build-windows
```

実行結果

```
$ make build-windows
echo "Build for windows10"
Build for windows10
GOOS=windows GOARCH=amd64 go build -o dist/windows/nemprice-win.exe app.go
echo "Done!"
Done!
```

`Windows 10`や`Ubuntu`用にビルドした実行ファイルを、アンドロイドスマホからコピーするには`scp`コマンドを使うのが一番簡単だとは思いますが、ちょっとだけコツがいるので、またの機会に紹介したいと思います


同様にMac用にビルドしたい場合は次のコマンドを実行してください。

```
$ make build-mac
```

実行結果

```
$ make build-mac
echo "Build for macOS(Darwin)"
Build for macOS(Darwin)
GOOS=darwin GOARCH=amd64 go build -o dist/macos/nemprice-mac app.go
```

## WEBアプリの起動

```
$ make run
```

またはビルドしてから実行する場合は、

```
$ make build
```

次にビルドされたアプリを起動します。

```
$ ./app
```

上記の手順を実行すると、`app`というWEBアプリが起動します。

## WEBアプリの動作確認

アンドロイドスマホ上のTermuxで直接、またはPCのターミナルからSSH接続して、WEBアプリを起動しておいてください。
起動方法は前述のとおりです。


動作確認の方法は、アンドロイドスマホ上のChromeアプリで確認する方法と、PC上のChromeアプリで確認する方法があります。

### アンドロイドスマホ上のChromeアプリで確認する方法

- Chromeブラウザアプリを起動します。
- http://localhost:3000/last_price にアクセスします。

次の結果が得られるはずです。

<img src="https://i.gyazo.com/dd14302368872c8a7349fdb887238ff2.png" alt="drawing" width="30%"/>


### PC上のChromeアプリで確認する方法


- Chromeブラウザアプリを起動します。
- http://アンドロイドスマホのIPアドレス:3000/last_price にアクセスします。

IPアドレスの調べ方は次のコマンドを実行してください。

```
$ ip -4 a | grep inet
```

実行結果

```
$ ip -4 a | grep inet
    inet 127.0.0.1/8 scope host lo
    inet 192.168.1.3/24 brd 192.168.1.255 scope global wlan0
```

この場合IPアドレスは`192.168.1.3'です。

ですので、PCのブラウザから`http://192.168.1.3:3000/last_price`にアクセスしてください。



次の結果が得られるはずです。

![](https://i.gyazo.com/44921c9ba402dc614f41b895dee44248.png)


WEBアプリのログはこんな感じで表示されます。

<img src="https://i.gyazo.com/923f838bdac5c21d43846d7f57a33393.png" alt="drawing" width="50%"/>

## ソースコード

`Tips`シリーズの記事では原則、ソースコードの説明はしません。
コピペや今回のように、GitHubからリポジトリをダウンロードしてきて、そのまま使える内容をご紹介していますので、ご了承ください。

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

国家試験でPythonを選択できるようになりましたが、馬鹿じゃないのか？と思います。時代が変わったといえ、スクリプト言語を国家試験レベルの対象にすべきではないと私は思ってます。

ということで、今後のTipsシリーズではPython言語も扱いますが、わかりやすい場合はGo言語で紹介したいと思います。

それにしても、アンドロイドスマホでこれだけの事ができるなんてすごくないですか？

確かに速度は遅いですが、やりたいこと（簡単な開発）はすべてできています。

環境をぶち壊しても、Termuxアプリを削除して、環境を作り直すだけです。

環境を作り直すのは若干面倒なので、自動で再構築するスクリプトを提供する予定ですのでお楽しみに。

## 関連情報へのリンク

- [techbureau/zaifapi: zaifのAPIを簡単にコール出来るようにしました。](https://github.com/techbureau/zaifapi)
- [gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.](https://github.com/)
- [ZaifAPI ドキュメント — Zaif api document v1.1.1 ドキュメント](https://techbureau-api-document.readthedocs.io/ja/latest/index.html)
- https://scrapbox.io/api/code/naoland/mypage/app.go

