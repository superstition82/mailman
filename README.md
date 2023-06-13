# mailman 🗞

## API 

## email

- [x] POST /api/email/send

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
    - 편의를 위해 실제 파일은 따로 보관하지 않음.

### 이슈

- redis, pulsar와 같은 외부 의존성을 줄이기 위해 태스크 작업을 미지원
  - 전부 프론트 단에서 API를 여러번 호출하도록 만듦.
