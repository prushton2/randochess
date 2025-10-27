import React from 'react'
import ReactDOM from 'react-dom/client'
import Join from './Join.tsx'
import Game from './Game.tsx'
import './index.css'
import {createBrowserRouter, RouterProvider} from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Join />,
  },
  {
   path: "/play",
   element: <Game />
  }
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
	<RouterProvider router={router} />
  </React.StrictMode>,
)
