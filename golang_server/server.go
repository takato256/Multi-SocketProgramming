package main

import (
	"fmt"
	"net"
)

func main() {
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// 関数が戻る時にコネクションを閉じる
	defer conn.Close()

	// データを受け取るためのバッファ
	buf := make([]byte, 1024)

	// コネクションが閉じるまでデータを読み取る
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Connection closed.")
			return
		}

		// 受信したバイト列を文字列に変換して表示
		msg := string(buf[:n])
		fmt.Printf("Received message: %s\n", msg)

		// クライアントにメッセージを返す
		resp := "Message received"
		_, err = conn.Write([]byte(resp))
		if err != nil {
			panic(err)
		}
		fmt.Printf("Sent message: %s\n", resp)
	}
}

