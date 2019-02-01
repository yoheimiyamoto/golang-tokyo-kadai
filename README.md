## 概要
golang.tokyo #21 課題
* [commpass](https://golangtokyo.connpass.com/event/113768/)
* [課題](https://docs.google.com/document/d/1jk8Ri6nogRdIqIaxCgqZ4MiS9Sq2roQ53Nipo7YQqeE/edit)

## 使い方
### 実行例
```
./tree -d {任意のディレクトリパス} -l {最大の階層の深さ} -e {対象の拡張子}
```

### コマンドライン引数
* `d`  
tree出力する対象のフォルダ。パラメータを省略した場合はカレントディレクトとなります。
* `l`  
出力する最大の階層の深さ。パラメータを省略した場合は最大の深さは100となります。 
* `e`  
対象のファイル拡張子のみを出力。パラメータを省略した場合は全拡張子が出力されます。

### 出力結果イメージ
```
.
├── 1.txt
├── dir1
│    ├── 1.txt
│    ├── 2.txt
│    └── 3.txt
└── dir2
     ├── 1.txt
     ├── 2.txt
     └── dir2-1
          ├── 1.txt
```


## 参考
https://github.com/gosagawa/golangtokyo21/
https://github.com/sh0e1/gotree
https://github.com/po3rin/gotree
https://github.com/mura-s/go-tree
https://github.com/pankona/ki