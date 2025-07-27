import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { get } from "../../utils/api";
import {
  Box,
  Heading,
  Text,
  Badge,
  VStack,
  HStack,
  Spinner,
  Flex,
  Image,
  Card,
  Separator,
  Stack,
} from "@chakra-ui/react";
import { Button } from "../ui/button";
import { FiEdit, FiArrowLeft } from "react-icons/fi";

interface Item {
  item_id: string;
  item_name: string;
  stock: boolean;
  description: string;
  created_at: string;
  updated_at: string;
  image_url?: string;
}

interface AdminItemDetailProps {
  itemId?: string;
  onClose?: () => void;
  onEdit?: (item: Item) => void;
}

const AdminItemDetail: React.FC<AdminItemDetailProps> = ({
  itemId,
  onClose,
  onEdit,
}) => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [item, setItem] = useState<Item | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchItem = async () => {
      const effectiveId = itemId || id;
      if (!effectiveId) return;

      try {
        setLoading(true);
        const response = await get(`/v1/admin/items/${effectiveId}`, true);

        if (response.ok) {
          const data = await response.json();
          setItem(data);
          setError(null);
        } else {
          setError("商品の取得に失敗しました");
        }
      } catch {
        setError("ネットワークエラーが発生しました");
      } finally {
        setLoading(false);
      }
    };

    fetchItem();
  }, [id, itemId]);

  if (loading) {
    return (
      <Flex justify="center" align="center" h="64">
        <Spinner size="xl" />
      </Flex>
    );
  }

  if (error || !item) {
    return (
      <Box bg="red.50" color="red.600" p={3} borderRadius="md">
        {error || "商品が見つかりません"}
      </Box>
    );
  }

  return (
    <Box>
      <Flex justify="space-between" align="center" mb={6}>
        <HStack gap={4}>
          <Button
            variant="ghost"
            onClick={() => {
              if (onClose) {
                onClose();
              } else {
                navigate("/admin/items");
              }
            }}
          >
            <FiArrowLeft /> 戻る
          </Button>
          <Heading size="lg">商品詳細</Heading>
        </HStack>
        <Button
          colorScheme="blue"
          onClick={() => {
            if (onEdit && item) {
              onEdit(item);
            } else {
              const effectiveId = itemId || id;
              navigate(`/admin/items/${effectiveId}/edit`);
            }
          }}
        >
          <FiEdit /> 編集
        </Button>
      </Flex>

      <Card.Root>
        <Card.Header>
          <HStack gap={4}>
            <Image
              src={
                item.image_url ||
                "https://images.unsplash.com/photo-1555041469-a586c61ea9bc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80"
              }
              alt={item.item_name}
              width="100px"
              height="100px"
              objectFit="cover"
              borderRadius="md"
            />
            <VStack align="start" gap={2}>
              <Heading size="md">{item.item_name}</Heading>
              <Badge
                colorScheme={item.stock ? "green" : "red"}
                variant="subtle"
                fontSize="md"
              >
                {item.stock ? "在庫あり" : "在庫なし"}
              </Badge>
            </VStack>
          </HStack>
        </Card.Header>

        <Separator />

        <Card.Body>
          <Stack gap={4}>
            <Box>
              <Text fontWeight="bold" mb={2}>
                商品説明
              </Text>
              <Text whiteSpace="pre-wrap">
                {item.description || "説明なし"}
              </Text>
            </Box>

            <Separator />

            <HStack gap={8}>
              <Box>
                <Text fontWeight="bold" fontSize="sm" color="gray.600">
                  作成日
                </Text>
                <Text>{new Date(item.created_at).toLocaleString("ja-JP")}</Text>
              </Box>
              <Box>
                <Text fontWeight="bold" fontSize="sm" color="gray.600">
                  更新日
                </Text>
                <Text>{new Date(item.updated_at).toLocaleString("ja-JP")}</Text>
              </Box>
            </HStack>
          </Stack>
        </Card.Body>
      </Card.Root>
    </Box>
  );
};

export default AdminItemDetail;
