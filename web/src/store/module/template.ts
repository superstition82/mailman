import * as api from "../../helpers/api";
import store, { useAppSelector } from "..";
import * as reducer from "../reducer/template";
import {
  createTemplate,
  deleteTemplate,
  upsertTemplates,
  setIsFetching,
} from "../reducer/template";

export const useTemplateStore = () => {
  const state = useAppSelector((state) => state.template);
  const fetchTemplateById = async (id: TemplateId) => {
    const { data } = (await api.getTemplateById(id)).data;
    return data;
  };

  return {
    state,
    fetchTemplateById,
    fetchTemplates: async (limit?: number, offset?: number) => {
      store.dispatch(setIsFetching(true));
      const templateFind = { limit, offset };
      const { data } = (await api.getTemplateList(templateFind)).data;
      store.dispatch(upsertTemplates(data));
      store.dispatch(setIsFetching(false));
      return data;
    },
    getTemplateById: async (templateId: TemplateId) => {
      for (const s of state.templates) {
        if (s.id === templateId) {
          return s;
        }
      }
      return await fetchTemplateById(templateId);
    },
    createTemplate: async (templateCreate: TemplateCreate) => {
      const { data } = (await api.createTemplate(templateCreate)).data;
      store.dispatch(createTemplate(data));
      return data;
    },
    deleteTemplateById: async (id: TemplateId) => {
      await api.deleteSender(id);
      store.dispatch(deleteTemplate(id));
    },
    deleteBulkTemplate: async (templateIds: TemplateId[]) => {
      await api.deleteBulkSender(templateIds);
      store.dispatch(reducer.deleteBulkTemplate(templateIds));
    },
    toggleSelectAll: (checked: boolean) => {
      store.dispatch(reducer.toggleSelectAll(checked));
    },
    toggleSelect: (id: SenderId) => {
      store.dispatch(reducer.toggleSelect(id));
    },
    isSelected: (id: SenderId) => {
      return state.selectedIds.includes(id);
    },
  };
};
