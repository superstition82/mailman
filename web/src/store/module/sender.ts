import * as api from "../../helpers/api";
import store, { useAppSelector } from "..";
import * as reducer from "../reducer/sender";
import {
  createSender,
  deleteSender,
  setIsFetching,
  upsertSenders,
} from "../reducer/sender";

export const useSenderStore = () => {
  const state = useAppSelector((state) => state.sender);
  const fetchSenderById = async (senderId: SenderId) => {
    const { data } = (await api.getSenderById(senderId)).data;
    return data;
  };

  return {
    state,
    fetchSenders: async (limit = 10, offset = 0) => {
      store.dispatch(setIsFetching(true));
      const senderFind: SenderFind = {
        limit,
        offset,
      };
      const { data } = (await api.getSenderList(senderFind)).data;
      store.dispatch(upsertSenders(data));
      store.dispatch(setIsFetching(false));
      return data;
    },
    fetchSenderById,
    getSenderById: async (senderId: SenderId) => {
      for (const s of state.senders) {
        if (s.id === senderId) {
          return s;
        }
      }
      return await fetchSenderById(senderId);
    },
    createSender: async (senderCreate: SenderCreate) => {
      const { data } = (await api.createSender(senderCreate)).data;
      store.dispatch(createSender(data));
      return data;
    },
    deleteSenderById: async (senderId: SenderId) => {
      await api.deleteSender(senderId);
      store.dispatch(deleteSender(senderId));
    },
    deleteBulkSender: async (senderIds: SenderId[]) => {
      await api.deleteBulkSender(senderIds);
      store.dispatch(reducer.deleteBulkSender(senderIds));
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
