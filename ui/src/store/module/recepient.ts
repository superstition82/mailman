import * as api from "../../helpers/api";
import store, { useAppSelector } from "../";
import * as reducer from "../reducer/recepient";
import {
  createRecepient,
  deleteRecepient,
  setIsFetching,
  upsertRecepients,
} from "../reducer/recepient";

export const useRecepientStore = () => {
  const state = useAppSelector((state) => state.recepient);
  const fetchRecepientById = async (recepientId: RecepientId) => {
    const { data } = (await api.getRecepientById(recepientId)).data;
    return data;
  };

  return {
    state,
    getState: () => {
      return store.getState().recepient;
    },
    fetchRecepients: async (limit = 10, offset = 0) => {
      store.dispatch(setIsFetching(true));
      const recepientFind: RecepientFind = {
        limit,
        offset,
      };
      const { data } = (await api.getRecepientList(recepientFind)).data;
      store.dispatch(upsertRecepients(data));
      store.dispatch(setIsFetching(false));

      return data;
    },
    fetchRecepientById,
    getRecepientById: async (recepientId: RecepientId) => {
      for (const s of state.recepients) {
        if (s.id === recepientId) {
          return s;
        }
      }
      return await fetchRecepientById(recepientId);
    },
    createRecepient: async (recepientCreate: RecepientCreate) => {
      const { data } = (await api.createRecepient(recepientCreate)).data;
      store.dispatch(createRecepient(data));
      return data;
    },
    deleteRecepientById: async (recepientId: RecepientId) => {
      await api.deleteRecepient(recepientId);
      store.dispatch(deleteRecepient(recepientId));
    },
    deleteBulkRecepient: async (recepientIds: RecepientId[]) => {
      await api.deleteBulkRecepient(recepientIds);
      store.dispatch(reducer.deleteBulkRecepient(recepientIds));
    },
    validate: async (recepientId: RecepientId) => {
      const { data } = (await api.validateRecepient(recepientId)).data;
      store.dispatch(reducer.patchRecepient(data));
      return data;
    },
    toggleSelectAll: (checked: boolean) => {
      store.dispatch(reducer.toggleSelectAll(checked));
    },
    toggleSelect: (id: RecepientId) => {
      store.dispatch(reducer.toggleSelect(id));
    },
    isSelected: (id: RecepientId) => {
      return state.selectedIds.includes(id);
    },
  };
};
