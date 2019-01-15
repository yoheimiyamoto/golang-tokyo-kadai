## 概要
golang.tokyo #21 課題
* [commpass](https://golangtokyo.connpass.com/event/113768/)
* [課題](https://docs.google.com/document/d/1jk8Ri6nogRdIqIaxCgqZ4MiS9Sq2roQ53Nipo7YQqeE/edit)

## 使い方
### 実行例
```
./tree --dir={任意のディレクトリパス} --L={最大の階層の深さ}
```

### コマンドライン引数
* `dir`  
tree出力する対象のフォルダ。パラメータを省略した場合はカレントディレクトとなります。
* `L`  
出力する最大の階層の深さ。パラメータを省略した場合は最大の深さは100となります。 

## 参考
https://qiita.com/tanksuzuki/items/7866768c36e13f09eedb
https://github.com/xlab/treeprint