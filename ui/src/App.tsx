import React, { Suspense } from "react";
import { RouterProvider } from "react-router-dom";
import Loading from "./pages/Loading";
import { router } from "./router";

export const App: React.FC = () => {
  return (
    <Suspense fallback={<Loading />}>
      <RouterProvider router={router} />
    </Suspense>
  );
};
