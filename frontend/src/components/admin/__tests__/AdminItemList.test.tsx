import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, waitFor, fireEvent } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import AdminItemList from "../AdminItemList";

const mockGet = vi.fn();
const mockDel = vi.fn();

vi.mock("../../../utils/api", () => ({
  get: mockGet,
  del: mockDel,
}));

const renderWithRouter = (ui: React.ReactElement) => {
  return render(<BrowserRouter>{ui}</BrowserRouter>);
};

const mockItems = [
  {
    item_id: "1",
    item_name: "テスト商品1",
    stock: true,
    description: "テスト説明1",
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    item_id: "2",
    item_name: "テスト商品2",
    stock: false,
    description: "テスト説明2",
    created_at: "2024-01-02T00:00:00Z",
    updated_at: "2024-01-02T00:00:00Z",
  },
];

describe("AdminItemList", () => {
  beforeEach(() => {
    vi.clearAllMocks();
    window.confirm = vi.fn();
    window.alert = vi.fn();
  });

  it("商品一覧が正しく表示される", async () => {
    mockGet.mockResolvedValue({
      ok: true,
      json: async () => mockItems,
    });

    renderWithRouter(<AdminItemList />);

    await waitFor(() => {
      expect(screen.getByText("テスト商品1")).toBeInTheDocument();
      expect(screen.getByText("テスト商品2")).toBeInTheDocument();
    });

    expect(screen.getByText("在庫あり")).toBeInTheDocument();
    expect(screen.getByText("在庫なし")).toBeInTheDocument();
  });

  it("データ取得エラー時にエラーメッセージが表示される", async () => {
    mockGet.mockResolvedValue({
      ok: false,
    });

    renderWithRouter(<AdminItemList />);

    await waitFor(() => {
      expect(screen.getByText("商品の取得に失敗しました")).toBeInTheDocument();
    });
  });

  it("削除確認ダイアログで確認後、商品が削除される", async () => {
    mockGet.mockResolvedValue({
      ok: true,
      json: async () => mockItems,
    });

    mockDel.mockResolvedValue({
      ok: true,
    });

    (
      window.confirm as vi.MockedFunction<typeof window.confirm>
    ).mockReturnValue(true);

    renderWithRouter(<AdminItemList />);

    await waitFor(() => {
      expect(screen.getByText("テスト商品1")).toBeInTheDocument();
    });

    const deleteButton = screen.getAllByText("削除")[0];
    fireEvent.click(deleteButton);

    await waitFor(() => {
      expect(mockDel).toHaveBeenCalledWith("/v1/admin/items/1", true);
    });
  });

  it("削除確認ダイアログでキャンセルした場合、削除されない", async () => {
    mockGet.mockResolvedValue({
      ok: true,
      json: async () => mockItems,
    });

    (
      window.confirm as vi.MockedFunction<typeof window.confirm>
    ).mockReturnValue(false);

    renderWithRouter(<AdminItemList />);

    await waitFor(() => {
      expect(screen.getByText("テスト商品1")).toBeInTheDocument();
    });

    const deleteButton = screen.getAllByText("削除")[0];
    fireEvent.click(deleteButton);

    expect(mockDel).not.toHaveBeenCalled();
  });
});
