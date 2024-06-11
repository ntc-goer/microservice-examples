Order Service — Tạo một Order ở trạng thái APPROVAL_PENDING. \
Consumer Service — Xác minh rằng người tiêu dùng có thể đặt đơn hàng. \
Kitchen Service — Xác thực chi tiết đơn hàng và tạo một Ticket ở trạng thái CREATE_PENDING.\
Accounting Service — Phê duyệt thẻ tín dụng của người tiêu dùng.\
Kitchen Service — Thay đổi trạng thái của Ticket thành AWAITING_ACCEPTANCE.\
Order Service — Thay đổi trạng thái của Order thành APPROVED.

Tech Stack
- viper
  + Load Config file
- wire
  + Dependency injection 
  - grpc 
  + Consider to git clone https://github.com/protocolbuffers/protobuf.git
- grpc-gateway
  + Expose HTTP connection to service
- consul
  + Service Registration/Discovery
  + Service Discovery
- Resty
- saga
- circuit-breaker