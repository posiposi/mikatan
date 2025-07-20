import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import AdminItemForm from "../AdminItemForm";

const mockPost = vi.fn();
const mockPut = vi.fn();
const mockGet = vi.fn();
const mockNavigate = vi.fn();

vi.mock("../../../utils/api", () => ({
  post: mockPost,
  put: mockPut,
  get: mockGet,
}));

vi.mock("react-router-dom", async () => {
  const actual = await vi.importActual("react-router-dom");
  return {
    ...actual,
    useNavigate: () => mockNavigate,
    useParams: () => ({ id: "1" }),
  };
});

const renderWithRouter = (ui: React.ReactElement, initialEntries = ["/"]) => {
  return render(
    <MemoryRouter initialEntries={initialEntries}>{ui}</MemoryRouter>
  );
};

describe("AdminItemForm", () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  describe("新規作成モード", () => {
    it("フォームが正しく表示される", () => {
      renderWithRouter(<AdminItemForm mode="create" />);

      expect(screen.getByText("商品登録")).toBeInTheDocument();
      expect(screen.getByLabelText("商品名 *")).toBeInTheDocument();
      expect(screen.getByLabelText("在庫あり")).toBeInTheDocument();
      expect(screen.getByLabelText("商品説明")).toBeInTheDocument();
      expect(screen.getByText("登録")).toBeInTheDocument();
    });

    it("必須項目が空の場合、エラーメッセージが表示される", async () => {
      renderWithRouter(<AdminItemForm mode="create" />);

      const submitButton = screen.getByText("登録");
      fireEvent.click(submitButton);

      await waitFor(() => {
        expect(
          screen.getByText("商品名を入力してください")
        ).toBeInTheDocument();
      });
    });

    it("正しいデータでフォーム送信が成功する", async () => {
      mockPost.mockResolvedValue({ ok: true });

      renderWithRouter(<AdminItemForm mode="create" />);

      const nameInput = screen.getByLabelText("商品名 *");
      const descriptionInput = screen.getByLabelText("商品説明");
      const submitButton = screen.getByText("登録");

      fireEvent.change(nameInput, { target: { value: "テスト商品" } });
      fireEvent.change(descriptionInput, { target: { value: "テスト説明" } });
      fireEvent.click(submitButton);

      await waitFor(() => {
        expect(mockPost).toHaveBeenCalledWith(
          "/v1/admin/items",
          {
            item_name: "テスト商品",
            stock: true,
            description: "テスト説明",
          },
          true
        );
        expect(mockNavigate).toHaveBeenCalledWith("/admin/items");
      });
    });
  });

  describe("編集モード", () => {
    const mockItem = {
      item_id: "1",
      item_name: "既存商品",
      stock: false,
      description: "既存説明",
    };

    it("既存データがフォームに読み込まれる", async () => {
      mockGet.mockResolvedValue({
        ok: true,
        json: async () => mockItem,
      });

      renderWithRouter(<AdminItemForm mode="edit" />);

      await waitFor(() => {
        expect(screen.getByDisplayValue("既存商品")).toBeInTheDocument();
        expect(screen.getByDisplayValue("既存説明")).toBeInTheDocument();
      });
    });

    it("更新が正しく実行される", async () => {
      mockGet.mockResolvedValue({
        ok: true,
        json: async () => mockItem,
      });
      mockPut.mockResolvedValue({ ok: true });

      renderWithRouter(<AdminItemForm mode="edit" />);

      await waitFor(() => {
        expect(screen.getByDisplayValue("既存商品")).toBeInTheDocument();
      });

      const nameInput = screen.getByDisplayValue("既存商品");
      const submitButton = screen.getByText("更新");

      fireEvent.change(nameInput, { target: { value: "更新された商品" } });
      fireEvent.click(submitButton);

      await waitFor(() => {
        expect(mockPut).toHaveBeenCalledWith(
          "/v1/admin/items/1",
          {
            item_name: "更新された商品",
            stock: false,
            description: "既存説明",
          },
          true
        );
        expect(mockNavigate).toHaveBeenCalledWith("/admin/items");
      });
    });
  });
});
