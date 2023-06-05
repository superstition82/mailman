import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { uniqBy } from "lodash-es";

type State = {
  senders: Sender[];
  selectedIds: SenderId[];
  isFetching: boolean;
};

const senderSlice = createSlice({
  name: "sender",
  initialState: {
    senders: [],
    selectedIds: [],
    isFetching: true,
  } as State,
  reducers: {
    createSender: (state, action: PayloadAction<Sender>) => {
      return {
        ...state,
        senders: state.senders.concat(action.payload),
      };
    },
    deleteSender: (state, action: PayloadAction<SenderId>) => {
      return {
        ...state,
        senders: state.senders.filter((sender) => {
          return sender.id !== action.payload;
        }),
      };
    },
    upsertSenders: (state, action: PayloadAction<Sender[]>) => {
      return {
        ...state,
        senders: uniqBy([...state.senders, ...action.payload], "id"),
      };
    },
    setIsFetching: (state, action: PayloadAction<boolean>) => {
      return {
        ...state,
        isFetching: action.payload,
      };
    },
    toggleSelectAll: (state, action: PayloadAction<boolean>) => {
      return {
        ...state,
        selectedIds: action.payload
          ? state.senders.map((sender) => sender.id)
          : [],
      };
    },
    toggleSelect: (state, action: PayloadAction<number>) => {
      if (state.selectedIds.includes(action.payload)) {
        return {
          ...state,
          selectedIds: state.selectedIds.filter((id) => id !== action.payload),
        };
      } else {
        return {
          ...state,
          selectedIds: [...state.selectedIds, action.payload],
        };
      }
    },
  },
});

export const {
  createSender,
  deleteSender,
  upsertSenders,
  setIsFetching,
  toggleSelect,
  toggleSelectAll,
} = senderSlice.actions;

export default senderSlice.reducer;
