type SenderId = number;

type Sender = {
  id: SenderId;

  createdTs: TimeStamp;
  updatedTs: TimeStamp;

  host: string;
  port: number;
  email: string;
  password: string;
};

type SenderCreate = {
  host: string;
  port: number;
  email: string;
  password: string;
};

type SenderFind = {
  offset?: number;
  limit?: number;
};
