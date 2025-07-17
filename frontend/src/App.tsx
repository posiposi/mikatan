import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Navigation from "./components/Navigation";
import TopPage from "./components/TopPage";
import SignupPage from "./components/SignupPage";

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <Navigation />
        <TopPage />
      </>
    ),
  },
  {
    path: "/signup",
    element: (
      <>
        <Navigation />
        <SignupPage />
      </>
    ),
  },
]);

function App() {
  return (
    <RouterProvider
      router={router}
      future={{
        v7_startTransition: true,
      }}
    />
  );
}

export default App;
