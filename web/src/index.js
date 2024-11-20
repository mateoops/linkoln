import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { RouterProvider, useParams, createBrowserRouter } from "react-router-dom";

const router = createBrowserRouter ([
  {
    path: "/",
    element: <App />
  },
  {
    path: "/:url",
    element: <RedirectToBackend />
  }
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

function RedirectToBackend() {
  console.debug("xxx")
  const { url } = useParams(); // Get the "xxx" part of the route
  const newUrl = `http://localhost:8080/short/${url}`;
  window.location.replace(newUrl); // Redirect to the external URL
  return null; // Render nothing
}

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
