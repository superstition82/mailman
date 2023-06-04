import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
  showHeader: boolean;
}

const layoutSlice = createSlice({
  name: "layout",
  initialState: {
    showHeader: false,
  } as State,
  reducers: {
    setHeaderStatus: (state, action: PayloadAction<boolean>) => {
      return {
        ...state,
        showHeader: action.payload,
      };
    },
  },
});

export const { setHeaderStatus } = layoutSlice.actions;

export default layoutSlice.reducer;
