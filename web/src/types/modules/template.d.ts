type TemplateId = number;

type Template = {
  id: TemplateId;
  createdTs: TimeStamp;
  updatedTs: TimeStamp;
  subject: string;
  body: string;
  resourceIdList: number[];
};

type TemplateCreate = {
  subject: string;
  body: string;
  resourceIdList?: number[];
};

type TemplatePatch = {
  id: number;
  subject?: string;
  body?: string;
  resourceIdList?: number[];
};

type TemplateFind = {
  offset?: number;
  limit?: number;
};

type TemplateBulkDelete = {
  templates: number[];
};
