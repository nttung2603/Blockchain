# Blockchain

Mở các session khác nhau trên termial (thay cho các Node trong mạng Blockchain) và chạy chương trình `Golang` bằng command: `go run main.go`

`open <port>`: Lệnh này cho phép hệ thống mở một cổng kết nối để chấp nhận các kết nối đến từ các node trong mạng. Điều này là cần thiết để thiết lập liên lạc và chia sẻ thông tin giữa các node trong mạng peer-to-peer.
`connect <address>`: Lệnh này thiết lập một kết nối tới một địa chỉ được chỉ định, tạo ra một môi trường kết nối giữa hai node trong mạng.
`peers`: Lệnh này trả về danh sách các đối tác hiện tại mà hệ thống đang kết nối đến.
`createblockchain`: Khi sử dụng lệnh này, hệ thống sẽ bắt đầu quá trình tạo mới một chuỗi khối mới, bao gồm cả tạo ra khối đầu tiên (genesis block).
`blockchain`: Lệnh này cung cấp một cái nhìn tổng quan về trạng thái hiện tại của blockchain, bao gồm số lượng block, thông tin chi tiết về từng block và mối liên kết giữa chúng.
`block <index>`: Cho phép người dùng xem thông tin chi tiết về một block cụ thể trong chuỗi khối, dựa trên chỉ số index.
`mine <tx_1>,<tx_2>,...,<tx_n>`: Lệnh thực hiện quá trình đào block mới trong blockchain, với những transaction được chọn.
`clone <pid>`: Lệnh này cho phép hệ thống sao chép thông tin từ blockchain của một node cụ thể, giúp đồng bộ hóa dữ liệu giữa các node trong mạng.
`exit`: Kết thúc phiên làm việc và đóng chương trình.