# Golang-web-crawler

User gửi 1 hoặc NHIỀU link VN express lên server, server sẽ đi tới trang đó và crawl data về
 
Yêu cầu chi tiết:
 

------- DONE ---------
- Khi crawl data sẽ lấy các thông tin sau về title, heading, paragraph
- Data crawl về phải lưu lại vào 1 thư mục dưới dạng csv hoặc json
- Mỗi khi crawl data xong: các thông kê các thông tin sau và lưu vào SQL database
  + Các thông kê như ở bài 1
  + Các url hình ảnh trong bài
  + Các dạng url khác
  + Các thông tin thống kê đoạn văn ở bài 1
- Có thể download cái file csv hoặc json trên về bằng api
- Có limit về số lượt link có thể crawl trong khoảng thời gian nhất định để k bị trang chủ ban IP (thời gian giữa các link dc crawl)
- Crawl nhiều tầng, nghĩa là nếu trong bài đó có các bài artical khác thì tiếp tục crawl các link đó, user có thể optional từ 1 đến 3 tầng link như vậy
- Giữ các tag html nào
------- DONE ---------


------- DOING ---------
- Cho các option khi crawl : 
  + highlight các từ giống nhau bằng html tag hoặc in đậm từ đó lên (dùng Levenshtein Distance (Edit Distance_))
------- DOING ---------

------- NOT YET ---------
- Back end Golang: Log, Unit test, concurrency (mỗi link sẽ có 1 job chạy riêng biệt), style code, handle lỗi

- Front end react, UI đơn giản và hợp lí (Optional)
  + Form gửi link crawl
  + Trang xem data của từng link (nếu crawl link đó bị lỗi thì lưu lỗi vào để khi xem biết nó bị gì)
------- NOT YET ---------