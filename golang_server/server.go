package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("-----Go Server-----")
	// ポート8000で待ち受ける
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for clients...")

	for {
		// クライアントが接続するまで待機
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client connected.")

		// クライアントとの接続を別々のgoroutineで処理
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// 関数が戻る時にコネクションを閉じる
	defer conn.Close()

	// データを受け取るためのバッファ
	buf := make([]byte, 1024)

	// データを読み取る
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading data:", err.Error())
		return
	}

	// 受信したバイト列を文字列に変換して表示
	msg := string(buf[:n])
	catMessage := convertToCatLanguage(msg)


	_, err = conn.Write([]byte(catMessage))
	if err != nil{
		fmt.Println("Error sending data:", err.Error())
		return
	}

	fmt.Println("Message sent to client:", catMessage)
}

func convertToCatLanguage(message string) string {
	// 猫語に変換するロジックを実装する

	var builder strings.Builder
	for _, char := range message {
		// 文字コードをバイトスライスに変換
		bytes := []byte(string(char))

		// バイトスライスの各バイトを2ビットずつ取り出して猫語に変換
		for _, b := range bytes {
			// バイトの上位6ビットを無視
			b = b & 0x03

			// 2ビットごとに変換
			switch b {
			case 0x00:
				builder.WriteString("ニャン")
			case 0x01:
				builder.WriteString("ミャン")
			case 0x02:
				builder.WriteString("ミャオン")
			case 0x03:
				builder.WriteString("ニャーン")
			}
		}
	}

	return builder.String()
}
