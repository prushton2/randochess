import React from 'react'
import ReactDOM from 'react-dom/client'
import Join from './Join.tsx'
import Game from './Game.tsx'
import Game2 from './Game2.tsx'
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
  },
  {
   path: "/test",
   element: <Game2 />
  }
]);

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
	<RouterProvider router={router} />
  </React.StrictMode>,
)
