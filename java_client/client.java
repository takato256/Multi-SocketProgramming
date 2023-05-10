import java.io.*;
import java.net.*;
import java.util.Scanner;

public class client {

    public static void main(String[] args) throws Exception {
	
	System.out.println("-----Java client-----");
	
        // サーバーのIPアドレスとポート番号を指定
        String serverAddress = "127.0.0.1";
        int port = 8000;

        // サーバーに接続するためのソケットを作成
        Socket socket = new Socket(serverAddress, port);

        // 入力ストリームと出力ストリームを取得
        BufferedReader input = new BufferedReader(new InputStreamReader(socket.getInputStream()));
        PrintWriter output = new PrintWriter(socket.getOutputStream(), true);

        // サーバーにデータを送信
        Scanner scanner = new Scanner(System.in);
        System.out.println("Input message:");
        String message = scanner.next();
        output.println(message);

        // サーバーからの応答を受信
        String response = input.readLine();
        System.out.println("Server response: " + response);

        // 接続を閉じる
        socket.close();
    }
}

