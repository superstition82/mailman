import { useMemo, useState } from "react";

export const useLoading = () => {
  const [state, setState] = useState({
    isLoading: false,
    count: 0,
    total: 0,
  });

  const progress = useMemo(() => {
    const { count, total } = state;
    return (count / total) * 100;
  }, [state]);

  return {
    ...state,
    progress,
    setStart: (total: number) => {
      setState({
        ...state,
        isLoading: true,
        count: 0,
        total,
      });
    },
    setNextTick: () => {
      setState((prev) => ({
        ...prev,
        isLoading: true,
        count: prev.count + 1,
      }));
    },
    setFinish: () => {
      setState({
        ...state,
        isLoading: false,
        count: 0,
        total: 0,
      });
    },
  };
};
