type ResourceId = number;

type Resource = {
  id: ResourceId;

  createdTs: TimeStamp;
  updatedTs: TimeStamp;

  filename: string;
  externalLink: string;
  type: string;
  size: string;
  publicId: string;

  linkedTemplateAmount: number;
};

type ResourceCreate = {
  filename: string;
  externalLink: string;
  type: string;
  downloadToLocal: boolean;
};

type ResourceFind = {
  offset?: number;
  limit?: number;
};
