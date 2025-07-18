import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Navigation from "./components/Navigation";
import TopPage from "./components/TopPage";
import SignupPage from "./components/SignupPage";
import LoginPage from "./components/LoginPage";
import { AuthProvider } from "./contexts/AuthContext";

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
  {
    path: "/login",
    element: (
      <>
        <Navigation />
        <LoginPage />
      </>
    ),
  },
]);

function App() {
  return (
    <AuthProvider>
      <RouterProvider
        router={router}
        future={{
          v7_startTransition: true,
        }}
      />
    </AuthProvider>
  );
}

export default App;
