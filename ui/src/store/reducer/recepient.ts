import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { uniqBy } from "lodash-es";

type State = {
  recepients: Recepient[];
  selectedIds: RecepientId[];
  isFetching: boolean;
};

const recepientSlice = createSlice({
  name: "recepient",
  initialState: {
    recepients: [],
    selectedIds: [],
    isFetching: true,
  } as State,
  reducers: {
    createRecepient: (state, action: PayloadAction<Recepient>) => {
      return {
        ...state,
        recepients: state.recepients.concat(action.payload),
      };
    },
    deleteRecepient: (state, action: PayloadAction<RecepientId>) => {
      return {
        ...state,
        recepients: state.recepients.filter((recepient) => {
          return recepient.id !== action.payload;
        }),
      };
    },
    deleteBulkRecepient: (state, action: PayloadAction<RecepientId[]>) => {
      return {
        ...state,
        recepients: state.recepients.filter((recepient) => {
          return !action.payload.includes(recepient.id);
        }),
      };
    },
    upsertRecepients: (state, action: PayloadAction<Recepient[]>) => {
      return {
        ...state,
        recepients: uniqBy([...state.recepients, ...action.payload], "id"),
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
          ? state.recepients.map((recepient) => recepient.id)
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
  createRecepient,
  deleteRecepient,
  deleteBulkRecepient,
  setIsFetching,
  toggleSelect,
  toggleSelectAll,
  upsertRecepients,
} = recepientSlice.actions;

export default recepientSlice.reducer;
