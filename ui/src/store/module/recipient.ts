import * as api from "../../helpers/api";
import store, { useAppSelector } from "..";
import * as reducer from "../reducer/recipient";
import {
  createRecipient,
  deleteRecipient,
  setIsFetching,
  upsertRecipients,
} from "../reducer/recipient";

export const useRecipientStore = () => {
  const state = useAppSelector((state) => state.recipient);
  const fetchRecipientById = async (recipientId: RecipientId) => {
    const { data } = (await api.getRecipientById(recipientId)).data;
    return data;
  };

  return {
    state,
    getState: () => {
      return store.getState().recipient;
    },
    fetchRecipients: async (limit = 10, offset = 0) => {
      store.dispatch(setIsFetching(true));
      const recipientFind: RecipientFind = {
        limit,
        offset,
      };
      const { data } = (await api.getRecipientList(recipientFind)).data;
      store.dispatch(upsertRecipients(data));
      store.dispatch(setIsFetching(false));

      return data;
    },
    fetchRecipientById,
    getRecipientById: async (recipientId: RecipientId) => {
      for (const s of state.recipients) {
        if (s.id === recipientId) {
          return s;
        }
      }
      return await fetchRecipientById(recipientId);
    },
    createRecipient: async (recipientCreate: RecipientCreate) => {
      const { data } = (await api.createRecipient(recipientCreate)).data;
      store.dispatch(createRecipient(data));
      return data;
    },
    deleteRecipientById: async (recipientId: RecipientId) => {
      await api.deleteRecipient(recipientId);
      store.dispatch(deleteRecipient(recipientId));
    },
    deleteBulkRecipient: async (recipientIds: RecipientId[]) => {
      await api.deleteBulkRecipient(recipientIds);
      store.dispatch(reducer.deleteBulkRecipient(recipientIds));
    },
    validate: async (recipientId: RecipientId) => {
      const { data } = (await api.validateRecipient(recipientId)).data;
      store.dispatch(reducer.patchRecipient(data));
      return data;
    },
    toggleSelectAll: (checked: boolean) => {
      store.dispatch(reducer.toggleSelectAll(checked));
    },
    toggleSelect: (id: RecipientId) => {
      store.dispatch(reducer.toggleSelect(id));
    },
    isSelected: (id: RecipientId) => {
      return state.selectedIds.includes(id);
    },
  };
};
