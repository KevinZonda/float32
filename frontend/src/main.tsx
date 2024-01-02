import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import {App} from "./App.tsx";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {About} from "./About.tsx";

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

