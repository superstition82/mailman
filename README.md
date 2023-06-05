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
