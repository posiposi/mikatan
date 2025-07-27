import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Navigation from "./components/Navigation";
import TopPage from "./components/TopPage";
import SignupPage from "./components/SignupPage";
import LoginPage from "./components/LoginPage";
import AdminDashboard from "./components/admin/AdminDashboard";
import AdminHome from "./components/admin/AdminHome";
import AdminItemList from "./components/admin/AdminItemList";
import AdminItemForm from "./components/admin/AdminItemForm";
import AdminItemDetail from "./components/admin/AdminItemDetail";
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
  {
    path: "/admin",
    element: <AdminDashboard />,
    children: [
      {
        index: true,
        element: <AdminHome />,
      },
      {
        path: "items",
        element: <AdminItemList />,
      },
      {
        path: "items/new",
        element: <AdminItemForm mode="create" />,
      },
      {
        path: "items/:id",
        element: <AdminItemDetail />,
      },
      {
        path: "items/:id/edit",
        element: <AdminItemForm mode="edit" />,
      },
    ],
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
