import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { post } from "../../utils/api";
import {
  Input,
  Textarea,
  Button,
  Box,
  Heading,
  Flex,
  Stack,
} from "@chakra-ui/react";
import { Field } from "../ui/field";
import { Checkbox } from "../ui/checkbox";

interface ItemFormData {
  item_name: string;
  stock: boolean;
  description: string;
}

interface AdminCreateItemProps {
  onSuccess?: () => void;
  onCancel?: () => void;
}

const AdminCreateItem: React.FC<AdminCreateItemProps> = ({
  onSuccess,
  onCancel,
}) => {
  const navigate = useNavigate();

  const [formData, setFormData] = useState<ItemFormData>({
    item_name: "",
    stock: true,
    description: "",
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!formData.item_name.trim()) {
      setError("商品名を入力してください");
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const response = await post("/v1/admin/items", formData, true);

      if (response.ok) {
        if (onSuccess) {
          onSuccess();
        } else {
          navigate("/admin/items");
        }
      } else {
        const errorData = await response.text();
        setError(errorData || "商品の登録に失敗しました");
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

  return (
    <Box>
      {!onCancel && (
        <Heading size="lg" mb={6}>
          商品登録
        </Heading>
      )}

      <Box
        bg={onCancel ? "transparent" : "gray.800"}
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
                登録
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

export default AdminCreateItem;
