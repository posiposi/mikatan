import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { post, put, get } from "../../utils/api";
import {
  Input,
  Textarea,
  Button,
  Box,
  Heading,
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
}

interface AdminItemFormProps {
  mode: "create" | "edit";
  itemId?: string;
  onSuccess?: () => void;
  onCancel?: () => void;
}

const AdminItemForm: React.FC<AdminItemFormProps> = ({
  mode,
  itemId,
  onSuccess,
  onCancel,
}) => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const effectiveId = itemId || id;

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
    if (mode === "edit" && effectiveId) {
      fetchItem(effectiveId);
    }
  }, [mode, effectiveId]);

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
        response = await put(`/v1/admin/items/${effectiveId}`, formData, true);
      }

      if (response.ok) {
        if (onSuccess) {
          onSuccess();
        } else {
          navigate("/admin/items");
        }
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
      <Flex justify="center" align="center" h="64">
        <Spinner size="xl" />
      </Flex>
    );
  }

  return (
    <Box>
      {!onCancel && (
        <Heading size="lg" mb={6}>
          {mode === "create" ? "商品登録" : "商品編集"}
        </Heading>
      )}

      <Box
        bg={onCancel ? "transparent" : "white"}
        shadow={onCancel ? "none" : "md"}
        borderRadius="lg"
        p={onCancel ? 0 : 6}
      >
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

            <Flex gap={4}>
              <Button type="submit" colorScheme="blue" loading={loading}>
                {mode === "create" ? "登録" : "更新"}
              </Button>
              <Button
                type="button"
                variant="outline"
                onClick={() => {
                  if (onCancel) {
                    onCancel();
                  } else {
                    navigate("/admin/items");
                  }
                }}
              >
                キャンセル
              </Button>
            </Flex>
          </Stack>
        </form>
      </Box>
    </Box>
  );
};

export default AdminItemForm;
