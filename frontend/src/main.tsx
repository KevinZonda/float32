import React from 'react'
import ReactDOM from 'react-dom/client'
import {App} from "./App.tsx";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {About} from "./About.tsx";

import 'tdesign-react/es/style/index.css'; // 少量公共样式
import './index.css'

const router = createBrowserRouter([
  {
    path: "/about",
    element: <About/>,
  },
  {
    path: "/search",
    element: <App />,
  },
  {
    path: "/",
    element: <App />,
  }
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)

