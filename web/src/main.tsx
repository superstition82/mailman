import ReactDOM from "react-dom/client";
import { CssVarsProvider } from "@mui/joy";
import { Provider } from "react-redux";

import { App } from "./App.tsx";
import store from "./store/index.ts";
import theme from "./theme/index.ts";

import "./css/global.css";
import "./css/tailwind.css";

const container = document.getElementById("root");
const root = ReactDOM.createRoot(container as HTMLElement);
root.render(
  <Provider store={store}>
    <CssVarsProvider defaultMode="light" theme={theme}>
      <App />
    </CssVarsProvider>
  </Provider>
);
