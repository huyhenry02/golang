# HƯỚNG DẪN QUY TRÌNH THIẾT KẾ & CHUẨN HÓA CƠ SỞ DỮ LIỆU
*(Áp dụng cho dự án đã có Figma)*

---

## I. Chuẩn bị & làm việc với Figma

1. Truy cập file **Figma** của dự án cần thiết kế cơ sở dữ liệu.
2. Chuyển Figma sang **chế độ Code** để nắm cấu trúc và dữ liệu liên quan.

---

## II. Chuẩn bị hệ thống tài liệu database

3. Kiểm tra project đã tồn tại thư mục `docs` và các tài liệu database hay chưa.
    - Nếu **chưa tồn tại**, thực hiện tạo mới theo các bước bên dưới.

4. Prompt dùng cho AI:
   > **Tạo cây thư mục `docs` và subfolder `docs/database` nếu chưa tồn tại.  
   > Trong thư mục `docs/database` tạo 3 file Markdown rỗng:
   > - `DatabaseCommand.md`
   > - `Database.md`
   > - `Collections.md`**

5. Cập nhật nội dung cho **03 file trên** theo tài liệu mẫu đã cung cấp.

6. Prompt dùng cho AI:
   > **Đọc và phân tích toàn bộ nội dung các file trong `docs/database` để nắm cấu trúc, quy ước và nguyên tắc thiết kế database hiện tại.**

---

## III. Phân tích & điều chỉnh database hiện có

7. Sau khi AI đã đọc hiểu tài liệu:
    - Yêu cầu **điều chỉnh database Supabase hiện tại** cho phù hợp với tài liệu đã phân tích.

8. Tạo thêm file:
    - `docs/database/ProjectCollections.md`

9. Mô tả **toàn bộ các bảng/collection đang có và còn thiếu của dự án** vào file  
   `docs/database/ProjectCollections.md`, theo **định dạng trong `docs/database/Collections.md`**.

---

## IV. Chuẩn hóa database bằng NotebookLM

10. Mở **NotebookLM** và tạo **01 notebook** cho dự án.

11. Thêm các nguồn tài liệu sau:
- https://raw.githubusercontent.com/vhvplatform/go-framework/refs/heads/main/docs/database/DatabaseCommand.md
- https://raw.githubusercontent.com/vhvplatform/go-framework/refs/heads/main/docs/database/Database.md
- https://raw.githubusercontent.com/vhvplatform/go-framework/refs/heads/main/docs/database/Collections.md
- File `ProjectCollections.md` đã tạo
- Các tài liệu liên quan khác của dự án (nếu có)
    
---

## V. Đánh giá & thiết kế chi tiết database

12. Dựa trên các phân tích, thực hiện **đánh giá và chuẩn hóa database chi tiết cho từng bảng/collection**,  
    theo **loại CSDL phù hợp với mục đích sử dụng**.

### Prompt chuẩn dùng cho NotebookLM

> **Thiết kế chi tiết bảng/collection `XXX` sử dụng hệ quản trị `YYY`, đáp ứng các phân tích ở trên.  
> Áp dụng UUID v7 cho khóa chính.  
> Trình bày dưới dạng bảng gồm các cột sau:**
> - Field (Tên trường)
> - Data Type (Kiểu dữ liệu)
> - Nullable (Có cho phép NULL không)
> - Default (Giá trị mặc định)
> - Constraints & Validation Logic (Ràng buộc & logic kiểm tra)
> - Description (Mô tả nghiệp vụ)

Trong đó:
- `XXX`: tên bảng / collection
- `YYY`: loại CSDL phù hợp

---

## VI. Nguyên tắc phân loại hệ quản trị CSDL

Hệ thống mới sử dụng **đa CSDL**, phân loại theo mục đích:

1. **Dữ liệu quan trọng, giao dịch**: YugabyteDB
2. **Dữ liệu linh hoạt cao**: MongoDB
3. **Log & thống kê**: ClickHouse
4. **Caching**: Redis (có thể nghiên cứu Dragonfly)
5. **Tìm kiếm toàn văn**: Elasticsearch
6. **Đồng bộ dữ liệu**: Debezium
7. **Graph Database**: FalkorDB
8. **Message Queue**: Kafka

---

## VII. Công việc bắt buộc sau khi hoàn tất thiết kế database

13. Sau khi thiết kế database hoàn tất, cần thực hiện **đầy đủ các bước sau**:

- Sao chép toàn bộ thiết kế database vào **bảng phân tích code Excel**  
  (Ví dụ:  
  https://docs.google.com/spreadsheets/d/1pKFvVOZqo5_Sttepv0DVUPugRCjAnLGIF2QpcYiCWhE/edit?gid=87378879#gid=87378879)

- Yêu cầu AI gửi lại:
    - Toàn bộ **DDL tạo bảng / collection**
    - Câu lệnh **tạo index**
    - Phân tách rõ theo **từng loại CSDL**

- Yêu cầu gửi lại **file `ProjectCollections.md` đã cập nhật hoàn chỉnh**

- Cập nhật lại nội dung database trong **Figma** theo file mới

- Yêu cầu Figma **đọc và hiểu lại file database mới**

- Truy cập **Supabase của dự án** để:
    - Thực thi các câu lệnh SQL đã thiết kế  
      *(Lưu ý: Figma không tự động chạy SQL)*

- Điều chỉnh lại từng phần chức năng theo database mới để **đảm bảo hệ thống hoạt động đúng**

---

## VIII. Ghi chú

- Đảm bảo đồng bộ giữa **Figma – Tài liệu – Database thực tế**
- Thiết kế phải đáp ứng:
    - Khả năng mở rộng
    - Khả năng truy vết
    - Dễ bảo trì
- Tuân thủ kiến trúc **đa CSDL** của hệ thống
