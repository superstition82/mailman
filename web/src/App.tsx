import React, { Suspense } from "react";
import { RouterProvider } from "react-router-dom";
import { router } from "./router";
import { Loading } from "./pages/Loading";

import "./styles/global.css";

export const App: React.FC = () => {
  return (
    <Suspense fallback={<Loading />}>
      <RouterProvider router={router} />
    </Suspense>
  );
};
