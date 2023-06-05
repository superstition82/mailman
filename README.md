# pocketmail ✉️

## API Docs

### sender

- /api/sender/
  - sender 생성, 목록

### recipient

- [x] /api/recipients/
  - [x] 생성
  - [x] 목록
- [x] /api/recipients/<recipient_id>/
  - 수신자 조회, 삭제
- [x] /api/recipients/<recipient_id>/verification
  - ID 리스트를 받아 검증 시작

### template

- /api/templates/
  - 템플릿 생성, 목록
- /api/templates/<template_id>/
  - 템플릿 조회, 수정, 삭제
  - 이미지, 파일 업로드 기능

### 특이사항

- 검증, 메일 발송 중에 변화를 계속 확인하면서 멈출 수 있어야하므로 클라이언트에서 계속 API 콜을 하는 식으로 만들기로 함.
