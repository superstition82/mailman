import { createBrowserRouter } from "react-router-dom";
import { lazy } from "react";

const Root = lazy(() => import("../layouts/Root"));
const Home = lazy(() => import("../pages/Home"));
const SenderManagement = lazy(() => import("../pages/SenderManagement"));
const RecepientManagement = lazy(() => import("../pages/RecepientManagement"));
const NotFound = lazy(() => import("../pages/NotFound"));

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "",
        element: <Home />,
      },
      {
        path: "/manage/sender",
        element: <SenderManagement />,
      },
      {
        path: "/manage/recepient",
        element: <RecepientManagement />,
      },
    ],
  },
  {
    path: "*",
    element: <NotFound />,
  },
]);
