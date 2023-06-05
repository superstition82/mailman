type RecepientId = number;

type Recepient = {
  id: RecepientId;

  createdTs: TimeStamp;
  updatedTs: TimeStamp;

  email: string;
  reachable: string;
};

type RecepientCreate = {
  email: string;
};

type RecepientFind = {
  offset?: number;
  limit?: number;
};
