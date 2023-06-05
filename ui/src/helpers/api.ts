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

export function getRecepientList(recepientFind?: RecepientFind) {
  return axios.get<ResponseObject<Recepient[]>>(`/api/recepient`, {
    params: {
      offset: recepientFind?.offset,
      limit: recepientFind?.limit,
    },
  });
}

export function getRecepientById(recepientId: RecepientId) {
  return axios.get<ResponseObject<Recepient>>(`/api/recepient/${recepientId}`);
}

export function createRecepient(recepientCreate: RecepientCreate) {
  return axios.post<ResponseObject<Recepient>>(
    `/api/recepient`,
    recepientCreate
  );
}

export function deleteRecepient(recepientId: RecepientId) {
  return axios.delete(`/api/recepient/${recepientId}`);
}

export function validateRecepient(recepientId: RecepientId) {
  return axios.post<ResponseObject<Recepient>>(
    `/api/recepient/${recepientId}/verification`
  );
}
