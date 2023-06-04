import React from "react";
import ReactDOM from "react-dom/client";
import { CssVarsProvider } from "@mui/joy";
import { Provider } from "react-redux";

import store from "./store";
import theme from "./theme";
import { App } from "./App.tsx";
import "./styles/tailwind.css";

const container = document.getElementById("root");
const root = ReactDOM.createRoot(container as HTMLElement);
root.render(
  <React.StrictMode>
    <Provider store={store}>
      <CssVarsProvider theme={theme}>
        <App />
      </CssVarsProvider>
    </Provider>
  </React.StrictMode>
);
