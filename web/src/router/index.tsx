import { createBrowserRouter } from "react-router-dom";
import { lazy } from "react";

const Root = lazy(() => import("../layouts/Root"));
const EmailDashboard = lazy(() => import("../pages/TemplateDashboard"));
const TemplateWrite = lazy(() => import("../pages/TemplateWrite"));
const SenderDashboard = lazy(() => import("../pages/SenderDashboard"));
const RecipientDashboard = lazy(() => import("../pages/RecipientDashboard"));
const ResourcesDashboard = lazy(() => import("../pages/ResourcesDashboard"));
const CampaignDashboard = lazy(() => import("../pages/CampaignDashboard"));
const NotFound = lazy(() => import("../pages/NotFound"));

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "",
        element: <EmailDashboard />,
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
      {
        path: "/campaign",
        element: <CampaignDashboard />,
      },
    ],
  },
  {
    path: "/write",
    element: <TemplateWrite />,
  },
  {
    path: "*",
    element: <NotFound />,
  },
]);
