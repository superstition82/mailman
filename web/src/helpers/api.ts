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

export function findRecipientList(recipientFind?: RecipientFind) {
  return axios.get<ResponseObject<Recipient[]>>(`/api/recipient`, {
    params: {
      offset: recipientFind?.offset,
      limit: recipientFind?.limit,
    },
  });
}

export function findRecipientById(recipientId: RecipientId) {
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

export function importRecipientFile(file: File) {
  const formData = new FormData();
  formData.append("file", file);

  return axios.post<ResponseObject<Recipient[]>>(
    `/api/recipient/file-import`,
    formData,
    {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    }
  );
}

export function exportRecipientFile() {
  return axios.get(`/api/recipient/file-export`, { responseType: "text" });
}

export function findTemplateList(templateFind?: TemplateFind) {
  return axios.get<ResponseObject<Template[]>>(`/api/template`, {
    params: {
      offset: templateFind?.offset,
      limit: templateFind?.limit,
    },
  });
}

export function findTemplateById(templateId: TemplateId) {
  return axios.get<ResponseObject<Template>>(`/api/template/${templateId}`);
}

export function createTemplate(templateCreate: TemplateCreate) {
  return axios.post<ResponseObject<Template>>(`/api/template`, templateCreate);
}

export function patchTemplate(templatePatch: TemplatePatch) {
  return axios.patch<ResponseObject<Template>>(
    `/api/template/${templatePatch.id}`,
    templatePatch
  );
}

export function deleteTemplate(templateId: TemplateId) {
  return axios.delete(`/api/template/${templateId}`);
}

export function deleteBulkTemplate(templateIds: TemplateId[]) {
  return axios.post(`/api/template/bulk-delete`, {
    templates: templateIds,
  });
}

export function getResourceList() {
  return axios.get<ResponseObject<Resource[]>>("/api/resource");
}

export function getResourceListWithLimit(resourceFind?: ResourceFind) {
  const queryList = [];
  if (resourceFind?.offset) {
    queryList.push(`offset=${resourceFind.offset}`);
  }
  if (resourceFind?.limit) {
    queryList.push(`limit=${resourceFind.limit}`);
  }
  return axios.get<ResponseObject<Resource[]>>(
    `/api/resource?${queryList.join("&")}`
  );
}

export function createResource(resourceCreate: ResourceCreate) {
  return axios.post<ResponseObject<Resource>>("/api/resource", resourceCreate);
}

export function createResourceWithBlob(formData: FormData) {
  return axios.post<ResponseObject<Resource>>("/api/resource/blob", formData);
}

export function deleteResourceById(id: ResourceId) {
  return axios.delete(`/api/resource/${id}`);
}

type EmailSend = {
  template: number;
  sender: number;
  recipients: number[];
  bcc: number[];
};

export function sendEmail({ template, sender, recipients, bcc }: EmailSend) {
  return axios.post<string>(`/api/email/send`, {
    template,
    sender,
    recipients,
    bcc,
  });
}
