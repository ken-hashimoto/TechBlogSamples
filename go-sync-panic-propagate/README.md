# Go Sync Panic Propagation Samples

このリポジトリは、golang.org/x/sync パッケージの v0.14.0 で追加された panic 伝播機能のサンプルコードを含んでいます。

## セットアップ

```bash
# 依存関係をダウンロード
go mod tidy
```

## サンプルコード

### 1. 基本的な errgroup.Group の使い方
```bash
go run 01_basic_errgroup/main.go
```

複数の HTTP リクエストを並行実行し、エラーハンドリングを行う基本的な例です。

### 2. 従来の挙動での panic
```bash
go run 02_traditional_panic/main.go
```

v0.14.0 以降では panic が伝播されるようになった例です。

### 3. Panic 伝播の例
```bash
go run 03_panic_propagation/main.go
```

複数の goroutine で panic が発生した場合の伝播動作を確認できます。

### 4. PanicError/PanicValue 型の例
```bash
go run 04_panic_types/main.go
```

error 型と非 error 型の panic がどのように処理されるかを確認できます。

### 5. runtime.Goexit の伝播
```bash
go run 05_goexit_propagation/main.go
```

runtime.Goexit() の伝播動作を確認できます。

### 6. テスト例
```bash
cd 06_test_examples
go test -v
```

panic 伝播機能のテストコード例です。

### 7. 堅牢なエラーハンドリング
```bash
go run 07_robust_handling/main.go
```

実際のアプリケーションで使用できる堅牢な panic ハンドリングの例です。
複数回実行すると、ランダムに異なる動作（成功/エラー/panic）を確認できます。

## 使用している依存関係

- golang.org/x/sync v0.14.0

## バージョン間の挙動比較

v0.13.0 と v0.14.0 の挙動の違いを確認するには：

```bash
# 自動比較スクリプトを実行
./compare_versions.sh
```

または手動で比較：

```bash
# v0.13.0の挙動を確認
sed -i 's/v0.14.0/v0.13.0/' go.mod && go mod tidy
go run 02_traditional_panic/main.go

# v0.14.0の挙動を確認
sed -i 's/v0.13.0/v0.14.0/' go.mod && go mod tidy
go run 02_traditional_panic/main.go
```

## 注意点

- このサンプルコードは golang.org/x/sync v0.14.0 以降の機能を使用しています
- v0.13.0 以前では panic が発生した場合にプログラム全体が停止します
- 実際の本番環境では適切なログ出力とモニタリングを実装してください 