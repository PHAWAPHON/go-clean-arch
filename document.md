# clean bob golang

## จะแบ่งออกเป็น 4 ส่วน
## Domain
## Repository
## REST (Interface Adapter / Handler)


## รายละเอียดแต่ละเลเยอร์

### 1. Domain
- **หน้าที่:** ประกาศโครงสร้างข้อมูล (Struct) และ Business Rules
- **สิ่งที่อยู่ใน Domain:**
  - **Entities/Models:** โครงสร้างข้อมูลหลัก เช่น Article, PDF, MergeRequest, SplitRequest, CompressRequest
  - **Errors & Constants:** handle error เช่น `ErrNotFound`, `ErrConflict`, `ErrInternalServerError`

### 2. Repository
- **หน้าที่:** ติดต่อกับแหล่งข้อมูลภายนอก (Database, File System, External API) 
- **สิ่งที่อยู่ใน Repository:**
  - **Interface Definition:** กำหนด Interface สำหรับการเข้าถึงข้อมูล เช่น `ArticleRepository`, `PDFRepository`
  - **Concrete Implementation:** การ implement interface เช่น MySQL driver, pdfcpu สำหรับจัดการไฟล์ PDF

### 3. REST (Interface Adapter / Handler)
- **หน้าที่:** รับ Request จากผู้ใช้ (ผ่าน HTTP) แปลงข้อมูลให้เข้ากับ Domain แล้วเรียกใช้ Service เพื่อดำเนินการตาม Business Logic จากนั้นแปลง Response กลับเป็นรูปแบบที่ผู้ใช้เข้าใจ (เช่น JSON)
- **สิ่งที่อยู่ใน REST:**
  - **Handlers/Controllers:** รับและแปลงข้อมูลจาก HTTP Request, ส่งกลับ HTTP Response
  - **Routing:** กำหนด endpoint สำหรับ API เช่น `/articles`, `/pdf/merge`
  - **Validation & Error Handling:** ตรวจสอบความถูกต้องของข้อมูลและจัดการกับ Error ที่เกิดขึ้น