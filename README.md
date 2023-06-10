# mailman 🗞

## API Docs

### sender

- [x] /api/sender/
  - [x] 생성
  - [x] 목록

### recipient

- [x] /api/recipient/
  - [x] 생성
  - [x] 목록
- [x] /api/recipient/file-import
  - [x] 파일 읽어서 추가
- [x] /api/recipient/file-export
  - [x] 검증된 수신자를 파일 형태로 추출
- [x] /api/recipient/<recipient_id>/
  - 수신자 조회, 삭제
- [x] /api/recipient/<recipient_id>/verification
  - ID 리스트를 받아 검증 시작

### template

- [x] /api/templates/
  - [x] 템플릿 생성, 목록
- [x] /api/templates/<template_id>/
  - [x] 템플릿 조회, 수정, 삭제
  - [x] 이미지, 파일 업로드 기능
    - 파일을 blob에서 읽어들여서 stream으로 쏴준다.
    - 실제 파일 자체를 가지고 있지는 않음.

### 이슈

- 결국 모든 첨부파일은 데이터베이스로 관리하기로 결정함.
  - 이메일을 보낼 때 파일을 전달해야하는데, 이는 임시파일을 만들어서 보내는 식으로 처리할 예정
    - SMTP에서 뭘 요구하는지 확인이 필요하긴 함.
- ## 에디터는 외부 라이브러리를 쓰기로 결정
