FROM scratch

# goreleaserが作成したバイナリをコピー
COPY orders-api-by-echo /

# バイナリを実行するためのエントリーポイント
ENTRYPOINT ["/orders-api-by-echo"]

# アプリケーションがリッスンするポートを指定
EXPOSE 8080
