import React, { useState, useEffect } from "react";
import { put, get } from "../../utils/api";
import {
  Input,
  Textarea,
  Button,
  Box,
  Flex,
  Spinner,
  Stack,
} from "@chakra-ui/react";
import { Field } from "../ui/field";
import { Checkbox } from "../ui/checkbox";

interface ItemFormData {
  item_name: string;
  stock: boolean;
  description: string;
  price_without_tax?: number;
  tax_rate?: number;
  currency?: string;
}

interface AdminEditItemProps {
  itemId: string;
  onSuccess?: () => void;
  onCancel?: () => void;
}

const AdminEditItem: React.FC<AdminEditItemProps> = ({
  itemId,
  onSuccess,
  onCancel,
}) => {
  const [formData, setFormData] = useState<ItemFormData>({
    item_name: "",
    stock: true,
    description: "",
    price_without_tax: undefined,
    tax_rate: 10,
    currency: "JPY",
  });

  const [loading, setLoading] = useState(false);
  const [loadingData, setLoadingData] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchItem = async () => {
    try {
      setLoadingData(true);
      const response = await get(`/v1/admin/items/${itemId}`, true);

      if (response.ok) {
        const item = await response.json();
        setFormData({
          item_name: item.item_name,
          stock: item.stock,
          description: item.description,
          price_without_tax: item.price_without_tax,
          tax_rate: item.tax_rate || 10,
          currency: item.currency || "JPY",
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
    if (itemId) {
      fetchItem();
    }
  }, [itemId]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!formData.item_name.trim()) {
      setError("商品名を入力してください");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const response = await put(`/v1/admin/items/${itemId}`, formData, true);

      if (response.ok) {
        if (onSuccess) {
          onSuccess();
        }
      } else {
        const errorData = await response.text();
        setError(errorData || "商品の更新に失敗しました");
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
    } else if (name === "price_without_tax") {
      setFormData((prev) => ({
        ...prev,
        [name]: value ? parseInt(value, 10) : undefined,
      }));
    } else if (name === "tax_rate") {
      setFormData((prev) => ({
        ...prev,
        [name]: value ? parseFloat(value) : undefined,
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
      <Flex justify="center" align="center" h="64">
        <Spinner size="xl" />
      </Flex>
    );
  }

  return (
    <Box>
      {error && (
        <Box bg="red.50" color="red.600" p={3} borderRadius="md" mb={4}>
          {error}
        </Box>
      )}

      <form onSubmit={handleSubmit}>
        <Stack gap={6}>
          <Field label="商品名" required>
            <Input
              name="item_name"
              value={formData.item_name}
              onChange={handleChange}
              placeholder="商品名を入力してください"
            />
          </Field>

          <Box>
            <Checkbox
              name="stock"
              checked={formData.stock}
              onCheckedChange={(e) => {
                setFormData((prev) => ({
                  ...prev,
                  stock: e.checked as boolean,
                }));
              }}
            >
              在庫あり
            </Checkbox>
          </Box>

          <Field label="商品説明">
            <Textarea
              name="description"
              value={formData.description}
              onChange={handleChange}
              placeholder="商品の説明を入力してください"
              rows={4}
            />
          </Field>

          <Field label="税抜き価格">
            <Input
              name="price_without_tax"
              type="number"
              value={formData.price_without_tax || ""}
              onChange={handleChange}
              placeholder="税抜き価格を入力してください"
            />
          </Field>

          <Field label="税率 (%)">
            <Input
              name="tax_rate"
              type="number"
              step="0.1"
              value={formData.tax_rate || 10}
              onChange={handleChange}
              placeholder="税率を入力してください"
            />
          </Field>

          <Field label="通貨">
            <Input
              name="currency"
              value={formData.currency || ""}
              onChange={handleChange}
              placeholder="通貨を入力してください"
            />
          </Field>

          <Flex gap={4}>
            <Button type="submit" colorScheme="blue" loading={loading}>
              更新
            </Button>
            <Button type="button" variant="outline" onClick={onCancel}>
              キャンセル
            </Button>
          </Flex>
        </Stack>
      </form>
    </Box>
  );
};

export default AdminEditItem;
