import { describe, it, expect, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import AdminDashboard from "../AdminDashboard";
import { AuthContext } from "../../../contexts/AuthContext";

const mockUseAuth = vi.fn();

vi.mock("../../../hooks/useAuth", () => ({
  useAuth: () => mockUseAuth(),
}));

const renderWithRouter = (
  ui: React.ReactElement,
  authValue: { isAdmin: boolean }
) => {
  return render(
    <AuthContext.Provider value={authValue}>
      <BrowserRouter>{ui}</BrowserRouter>
    </AuthContext.Provider>
  );
};

describe("AdminDashboard", () => {
  it("管理者でない場合、アクセス権限エラーが表示される", () => {
    mockUseAuth.mockReturnValue({ isAdmin: false });

    renderWithRouter(<AdminDashboard />, { isAdmin: false });

    expect(screen.getByText("アクセス権限がありません")).toBeInTheDocument();
    expect(screen.getByText("管理者権限が必要です。")).toBeInTheDocument();
  });

  it("管理者の場合、サイドバーが表示される", () => {
    mockUseAuth.mockReturnValue({ isAdmin: true });

    renderWithRouter(<AdminDashboard />, { isAdmin: true });

    expect(screen.getByText("管理画面")).toBeInTheDocument();
    expect(screen.getByText("ダッシュボード")).toBeInTheDocument();
    expect(screen.getByText("商品一覧")).toBeInTheDocument();
    expect(screen.getByText("商品登録")).toBeInTheDocument();
  });
});
