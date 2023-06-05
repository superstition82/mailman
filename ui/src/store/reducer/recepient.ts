import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { uniqBy } from "lodash-es";

type State = {
  recipients: Recipient[];
  selectedIds: RecipientId[];
  isFetching: boolean;
};

const recipientSlice = createSlice({
  name: "recipient",
  initialState: {
    recipients: [],
    selectedIds: [],
    isFetching: true,
  } as State,
  reducers: {
    createRecipient: (state, action: PayloadAction<Recipient>) => {
      return {
        ...state,
        recipients: state.recipients.concat(action.payload),
      };
    },
    deleteRecipient: (state, action: PayloadAction<RecipientId>) => {
      return {
        ...state,
        recipients: state.recipients.filter((recipient) => {
          return recipient.id !== action.payload;
        }),
      };
    },
    deleteBulkRecipient: (state, action: PayloadAction<RecipientId[]>) => {
      return {
        ...state,
        recipients: state.recipients.filter((recipient) => {
          return !action.payload.includes(recipient.id);
        }),
      };
    },
    upsertRecipients: (state, action: PayloadAction<Recipient[]>) => {
      return {
        ...state,
        recipients: uniqBy([...state.recipients, ...action.payload], "id"),
      };
    },
    patchRecipient: (state, action: PayloadAction<Partial<Recipient>>) => {
      return {
        ...state,
        recipients: state.recipients.map((recipient) => {
          if (recipient.id === action.payload.id) {
            return {
              ...recipient,
              ...action.payload,
            };
          } else {
            return recipient;
          }
        }),
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
          ? state.recipients.map((recipient) => recipient.id)
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
  createRecipient,
  deleteRecipient,
  deleteBulkRecipient,
  patchRecipient,
  setIsFetching,
  toggleSelect,
  toggleSelectAll,
  upsertRecipients,
} = recipientSlice.actions;

export default recipientSlice.reducer;
