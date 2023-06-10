import { createBrowserRouter } from "react-router-dom";
import { lazy } from "react";

const Root = lazy(() => import("../layouts/Root"));
const Home = lazy(() => import("../pages/TemplateDashboard"));
const SenderDashboard = lazy(() => import("../pages/SenderDashboard"));
const RecipientDashboard = lazy(() => import("../pages/RecipientDashboard"));
const ResourcesDashboard = lazy(() => import("../pages/ResourcesDashboard"));
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
        path: "/sender",
        element: <SenderDashboard />,
      },
      {
        path: "/recipient",
        element: <RecipientDashboard />,
      },
      {
        path: "/resource",
        element: <ResourcesDashboard />,
      },
    ],
  },
  {
    path: "*",
    element: <NotFound />,
  },
]);
