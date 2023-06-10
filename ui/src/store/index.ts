import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import layoutReducer from "./reducer/layout";
import senderReducer from "./reducer/sender";
import recipientReducer from "./reducer/recipient";
import templateReducer from "./reducer/template";

const store = configureStore({
  reducer: {
    layout: layoutReducer,
    sender: senderReducer,
    recipient: recipientReducer,
    template: templateReducer,
  },
});

type AppState = ReturnType<typeof store.getState>;
type AppDispatch = typeof store.dispatch;

export const useAppSelector: TypedUseSelectorHook<AppState> = useSelector;
export const useAppDispatch: () => AppDispatch = useDispatch;

export default store;
