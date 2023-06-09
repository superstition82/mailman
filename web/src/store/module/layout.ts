import store, { useAppSelector } from "..";
import { setHeaderStatus } from "../reducer/layout";

export const useLayoutStore = () => {
  const state = useAppSelector((state) => state.layout);
  return {
    state,
    getState: () => {
      return store.getState();
    },
    setHeaderStatus: (showHeader: boolean) => {
      store.dispatch(setHeaderStatus(showHeader));
    },
  };
};
