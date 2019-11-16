数据转载应声明原创地址、如要作为商用应该向原作者商议后
# golang中文官网部分数据爬取
### 因为golang天然支持并发、所以在这个程序中无论是存储还是爬取都使用了多线程
在 bin 文件夹中已经打包了三个平台(windows、macOs、linux)的可执行文件
获取数据成功后会在当前运行的目录下新建一个 data.json 文件


## 命令行解释
-t 是开启数据存储(先要爬取才能够存储)
-s 启始页(1)
-e 结束页(2)

## 使用方法
### 1、爬取数据
./main -s 1 -e 205(因为目前go中文网主题是205页)
![Text image](https://note.youdao.com/yws/api/personal/file/WEB2ba8f49739cdcb5b874ca6e7f5ea9958?method=download&shareKey=203d1afca223b0e8e9930b4eb734a856)


### 2、存储数据
./main -t(注意数据库）
![Text image](https://note.youdao.com/yws/api/personal/file/WEB20437591714b663e4ab525efb5e2ac15?method=download&shareKey=c042ca04524105fea3b2aa622b7df2e9)
