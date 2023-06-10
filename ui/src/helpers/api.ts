import axios from "axios";

type ResponseObject<T> = {
  data: T;
  error?: string;
  message?: string;
};

export function getSenderList(senderFind?: SenderFind) {
  return axios.get<ResponseObject<Sender[]>>(`/api/sender`, {
    params: {
      offset: senderFind?.offset,
      limit: senderFind?.limit,
    },
  });
}

export function getSenderById(senderId: SenderId) {
  return axios.get<ResponseObject<Sender>>(`/api/sender/${senderId}`);
}

export function createSender(senderCreate: SenderCreate) {
  return axios.post<ResponseObject<Sender>>(`/api/sender`, senderCreate);
}

export function deleteSender(senderId: SenderId) {
  return axios.delete(`/api/sender/${senderId}`);
}

export function deleteBulkSender(senderIds: SenderId[]) {
  return axios.post(`/api/sender/bulk-delete`, {
    senders: senderIds,
  });
}

export function getRecipientList(recipientFind?: RecipientFind) {
  return axios.get<ResponseObject<Recipient[]>>(`/api/recipient`, {
    params: {
      offset: recipientFind?.offset,
      limit: recipientFind?.limit,
    },
  });
}

export function getRecipientById(recipientId: RecipientId) {
  return axios.get<ResponseObject<Recipient>>(`/api/recipient/${recipientId}`);
}

export function createRecipient(recipientCreate: RecipientCreate) {
  return axios.post<ResponseObject<Recipient>>(
    `/api/recipient`,
    recipientCreate
  );
}

export function deleteRecipient(recipientId: RecipientId) {
  return axios.delete(`/api/recipient/${recipientId}`);
}

export function deleteBulkRecipient(recipientIds: RecipientId[]) {
  return axios.post(`/api/recipient/bulk-delete`, {
    recipients: recipientIds,
  });
}

export function validateRecipient(recipientId: RecipientId) {
  return axios.post<ResponseObject<Recipient>>(
    `/api/recipient/${recipientId}/verification`
  );
}

export function getTemplateList(templateFind?: TemplateFind) {
  return axios.get<ResponseObject<Template[]>>(`/api/template`, {
    params: {
      offset: templateFind?.offset,
      limit: templateFind?.limit,
    },
  });
}

export function getTemplateById(templateId: TemplateId) {
  return axios.get<ResponseObject<Template>>(`/api/template/${templateId}`);
}

export function createTemplate(templateCreate: TemplateCreate) {
  return axios.post<ResponseObject<Template>>(`/api/template`, templateCreate);
}

export function deleteTemplate(templateId: TemplateId) {
  return axios.delete(`/api/template/${templateId}`);
}

export function deleteBulkTemplate(templateIds: TemplateId[]) {
  return axios.post(`/api/template/bulk-delete`, {
    templates: templateIds,
  });
}
