type TemplateId = number;

type Template = {
  id: TemplateId;
  createdTs: TimeStamp;
  updatedTs: TimeStamp;
  subject: string;
  body: string;
};

type TemplateCreate = {
  subject: string;
  body: string;
};

type TemplatePatch = {
  id: number;
  subject?: string;
  body?: string;
};

type TemplateFind = {
  offset?: number;
  limit?: number;
};

type TemplateBulkDelete = {
  templates: number[];
};
