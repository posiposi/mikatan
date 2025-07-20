import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { post, put, get } from "../../utils/api";

interface ItemFormData {
  item_name: string;
  stock: boolean;
  description: string;
}

interface AdminItemFormProps {
  mode: "create" | "edit";
}

const AdminItemForm: React.FC<AdminItemFormProps> = ({ mode }) => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();

  const [formData, setFormData] = useState<ItemFormData>({
    item_name: "",
    stock: true,
    description: "",
  });

  const [loading, setLoading] = useState(false);
  const [loadingData, setLoadingData] = useState(mode === "edit");
  const [error, setError] = useState<string | null>(null);

  const fetchItem = async (itemId: string) => {
    try {
      setLoadingData(true);
      const response = await get(`/v1/admin/items/${itemId}`, true);

      if (response.ok) {
        const item = await response.json();
        setFormData({
          item_name: item.item_name,
          stock: item.stock,
          description: item.description,
        });
        setError(null);
      } else {
        setError("商品の取得に失敗しました");
      }
    } catch {
      setError("ネットワークエラーが発生しました");
    } finally {
      setLoadingData(false);
    }
  };

  useEffect(() => {
    if (mode === "edit" && id) {
      fetchItem(id);
    }
  }, [mode, id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!formData.item_name.trim()) {
      setError("商品名を入力してください");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      let response;

      if (mode === "create") {
        response = await post("/v1/admin/items", formData, true);
      } else {
        response = await put(`/v1/admin/items/${id}`, formData, true);
      }

      if (response.ok) {
        navigate("/admin/items");
      } else {
        const errorData = await response.text();
        setError(
          errorData ||
            `商品の${mode === "create" ? "登録" : "更新"}に失敗しました`
        );
      }
    } catch {
      setError("ネットワークエラーが発生しました");
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value, type } = e.target;

    if (type === "checkbox") {
      const checkbox = e.target as HTMLInputElement;
      setFormData((prev) => ({
        ...prev,
        [name]: checkbox.checked,
      }));
    } else {
      setFormData((prev) => ({
        ...prev,
        [name]: value,
      }));
    }
  };

  if (loadingData) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="text-lg">データを読み込み中...</div>
      </div>
    );
  }

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">
        {mode === "create" ? "商品登録" : "商品編集"}
      </h1>

      <div className="bg-white shadow rounded-lg p-6">
        {error && (
          <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label
              htmlFor="item_name"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              商品名 *
            </label>
            <input
              type="text"
              id="item_name"
              name="item_name"
              value={formData.item_name}
              onChange={handleChange}
              required
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="商品名を入力してください"
            />
          </div>

          <div>
            <label className="flex items-center">
              <input
                type="checkbox"
                name="stock"
                checked={formData.stock}
                onChange={handleChange}
                className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500"
              />
              <span className="ml-2 text-sm font-medium text-gray-700">
                在庫あり
              </span>
            </label>
          </div>

          <div>
            <label
              htmlFor="description"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              商品説明
            </label>
            <textarea
              id="description"
              name="description"
              value={formData.description}
              onChange={handleChange}
              rows={4}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="商品の説明を入力してください"
            />
          </div>

          <div className="flex gap-4">
            <button
              type="submit"
              disabled={loading}
              className="bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700 disabled:opacity-50 transition-colors"
            >
              {loading ? "処理中..." : mode === "create" ? "登録" : "更新"}
            </button>
            <button
              type="button"
              onClick={() => navigate("/admin/items")}
              className="bg-gray-300 text-gray-700 px-6 py-2 rounded hover:bg-gray-400 transition-colors"
            >
              キャンセル
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AdminItemForm;
