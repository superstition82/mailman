import { createBrowserRouter } from "react-router-dom";
import { lazy } from "react";

const Root = lazy(() => import("../layouts/Root"));
const Home = lazy(() => import("../pages/TemplateManagement"));
const SenderManagement = lazy(() => import("../pages/SenderManagement"));
const RecipientManagement = lazy(() => import("../pages/RecipientManagement"));
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
        path: "/manage/recipient",
        element: <RecipientManagement />,
      },
    ],
  },
  {
    path: "*",
    element: <NotFound />,
  },
]);
