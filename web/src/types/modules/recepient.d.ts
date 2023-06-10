type RecipientId = number;

type Recipient = {
  id: RecipientId;

  createdTs: TimeStamp;
  updatedTs: TimeStamp;

  email: string;
  reachable: string;
};

type RecipientCreate = {
  email: string;
};

type RecipientFind = {
  offset?: number;
  limit?: number;
};

type RecepeintBulkDelete = {
  recipients: number[];
};
