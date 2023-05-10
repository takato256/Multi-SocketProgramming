use std::io;
use std::io::prelude::*;
use std::net::TcpStream;

fn main() -> std::io::Result<()> {
    println!("-----Rust Client-----");

    let mut stream = TcpStream::connect("127.0.0.1:8000")?;

    // メッセージ文字列をバイト配列に変換して、サーバーに送信
    println!("Input message:");
    let mut msg = String::new();
    io::stdin().read_line(&mut msg)
    	       .expect("Failed to read");
    stream.write(msg.as_bytes())?;
    println!("Sent message: {}", msg);

    // 受信データを格納するバッファ
    let mut buf = [0; 1024];

    // コネクションを閉じるまでデータを読み取る
    loop {
        let n = stream.read(&mut buf)?;
        if n == 0 {
            break;
        }

        // 受信したバイト列を文字列に変換して表示
        let resp = String::from_utf8_lossy(&buf[..n]);
        println!("Received server message: {}", resp);
    }

    Ok(())
}

