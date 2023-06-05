# mails

## API Docs

### sender

- [x] /api/sender/
  - [x] 생성
  - [x] 목록

### recipient

- [x] /api/recipients/
  - [x] 생성
  - [x] 목록
- [ ] /api/recepient/file-import
  - [ ] 파일 읽어서 추가
- [ ] /api/recepient/file-export
  - [ ] 검증된 수신자를 파일 형태로 추출
- [x] /api/recipients/<recipient_id>/
  - 수신자 조회, 삭제
- [x] /api/recipients/<recipient_id>/verification
  - ID 리스트를 받아 검증 시작

### template

- [ ] /api/templates/
  - [ ] 템플릿 생성, 목록
- [ ] /api/templates/<template_id>/
  - 템플릿 조회, 수정, 삭제
  - 이미지, 파일 업로드 기능

### 이슈

- 검증, 메일 발송 중에 변화를 계속 확인하면서 멈출 수 있어야하므로 클라이언트에서 계속 API 요청을 하는 식으로 만들기로 함.
- 템플릿 삭제했을 때 해당 템플릿에 포함된 파일들도 삭제되도록 설계
  - template_file 테이블을 생성한 뒤, 삭제될 때 path에 있는 파일을 삭제하는 로직 추가
  - 근데, 본문에서 지웠을 때 dangling 되는 경우는 해결할 수 없음..
    - 이건, 템플릿이 지워질 때 처리될 거라고 기대
- 이메일에 이미지를 파일을 통해 첨부하려면 base64로 인코딩한 후 blob으로 쏘는 형태와 cid를 이용한 방법이 있음.
  - 하지만, base64 방식으로 전달 시 대부분의 메일 클라이언트에서 제대로 처리해주지 않음.
  - 따라서 cid 방법을 이용해야하는데, 저장되는 본문과는 다르게 수정을 거쳐 메일을 전송해야 함.
    - regex를 사용하면 일반적인 경우에서는 해결할 수 있었음.
    - 더 좋은 방법이 있는지는 정말 고민되는 부분. 다른 구현체를 찾을 수 있으면 좋겠다.
