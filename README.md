# 留言小站 Comment Site
- Written in Golang（go version go1.18.1 darwin/amd64）
- Framework: gin
## How to run
移動到此作業的根⽬錄下
```shell
git clone https://github.com/wama-tw/commentSite.git
cd commentSite
```
執⾏指令
```shell
go run src/main.go
```
然後在喜歡的瀏覽器打開 http://localhost:8080

## Immplement details
- TL;DR
  - HTTP GET / POST
  - multi-threaded (Goroutine)
  - unbuffered channel
  - Parallel Merge Sort

⽤到 http 中的 GET 和 POST（GET "/posts" 取得顯⽰所有貼⽂的畫⾯，POST "/posts/create" 來傳送新⽂章的資料等），每⼀次的 request 都是⼀個新的 thread，因此就算同時有很多使⽤者，也不影響正常使⽤。
另外，除了 http server，在顯⽰⽂章的地⽅我也⽤了 Parallel Merge Sort 來排序⽂章顯⽰的順序（最新的會顯⽰在最上⾯）。在寫 Merge Sort 的時候，我是⽤ unbuffered channel，除了在 thread（Goroutine）之間傳遞資料之外，也因為 unbuffered channel 接收時會等待，所以也能順便同步 Merge Sort 中被切開的兩邊，等兩邊都被排好回傳的時候，才會把兩邊合併（merge）。最後也是⽤ channel 把結果傳回去。
