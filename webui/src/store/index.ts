import { configureStore } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import layoutReducer from "./reducer/layout";

const store = configureStore({
  reducer: {
    layout: layoutReducer,
  },
});

type AppState = ReturnType<typeof store.getState>;
type AppDispatch = typeof store.dispatch;

export const useAppSelector: TypedUseSelectorHook<AppState> = useSelector;
export const useAppDispatch: () => AppDispatch = useDispatch;

export default store;
