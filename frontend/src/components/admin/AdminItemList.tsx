import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { get, del } from "../../utils/api";
import {
  Table,
  Badge,
  Box,
  Flex,
  Heading,
  Spinner,
  Text,
  HStack,
  IconButton,
  VStack,
  PaginationItems,
  PaginationNextTrigger,
  PaginationPrevTrigger,
  PaginationRoot,
} from "@chakra-ui/react";
import { Button } from "../ui/button";
import {
  DialogRoot,
  DialogContent,
  DialogHeader,
  DialogBody,
  DialogFooter,
  DialogTitle,
  DialogActionTrigger,
  DialogCloseTrigger,
} from "../ui/dialog";
import { FiEdit, FiTrash2, FiPlus } from "react-icons/fi";
import AdminItemForm from "./AdminItemForm";
import AdminItemDetail from "./AdminItemDetail";

interface Item {
  item_id: string;
  item_name: string;
  stock: boolean;
  description: string;
  created_at: string;
  updated_at: string;
  image_url?: string;
}

const AdminItemList: React.FC = () => {
  const navigate = useNavigate();
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedItem, setSelectedItem] = useState<Item | null>(null);
  const [deleteItemId, setDeleteItemId] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage] = useState(10);

  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDeleteAlertOpen, setIsDeleteAlertOpen] = useState(false);
  const [isDetailModalOpen, setIsDetailModalOpen] = useState(false);

  const fetchItems = async () => {
    try {
      setLoading(true);
      const response = await get("/v1/admin/items", true);

      if (response.ok) {
        const data = await response.json();
        setItems(data);
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

  const handleEdit = (item: Item) => {
    setSelectedItem(item);
    setIsEditModalOpen(true);
  };

  const handleDeleteClick = (itemId: string) => {
    setDeleteItemId(itemId);
    setIsDeleteAlertOpen(true);
  };

  const handleDelete = async () => {
    if (!deleteItemId) return;

    try {
      const response = await del(`/v1/admin/items/${deleteItemId}`, true);

      if (response.ok) {
        setItems(items.filter((item) => item.item_id !== deleteItemId));
        setIsDeleteAlertOpen(false);
      } else {
        alert("削除に失敗しました");
      }
    } catch {
      alert("削除中にエラーが発生しました");
    }
  };

  const handleRowClick = (item: Item) => {
    setSelectedItem(item);
    setIsDetailModalOpen(true);
  };

  // ページネーション計算
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = items.slice(indexOfFirstItem, indexOfLastItem);
  const totalPages = Math.ceil(items.length / itemsPerPage);

  useEffect(() => {
    fetchItems();
  }, []);

  if (loading) {
    return (
      <Flex justify="center" align="center" h="64">
        <Spinner size="xl" />
      </Flex>
    );
  }

  if (error) {
    return (
      <Box bg="red.50" color="red.600" p={3} borderRadius="md">
        {error}
      </Box>
    );
  }

  return (
    <Box>
      <Flex justify="space-between" align="center" mb={6}>
        <Heading size="lg">商品一覧</Heading>
        <Button
          colorScheme="green"
          onClick={() => navigate("/admin/items/new")}
        >
          <FiPlus /> 新規登録
        </Button>
      </Flex>

      <VStack gap={6} align="stretch">
        <Box bg="gray.800" shadow="md" borderRadius="lg" overflow="hidden">
          <Table.Root
            variant="outline"
            size="sm"
            css={{
              "& thead tr": {
                backgroundColor: "var(--chakra-colors-gray-50)",
              },
            }}
          >
            <Table.Body>
              {currentItems.map((item) => (
                <Table.Row
                  key={item.item_id}
                  _hover={{
                    bg: "gray.50",
                    cursor: "pointer",
                    transform: "translateY(-1px)",
                    boxShadow: "sm",
                  }}
                  transition="all 0.2s"
                  onClick={() => handleRowClick(item)}
                >
                  <Table.Cell>
                    <Text fontWeight="medium">{item.item_name}</Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Badge
                      colorScheme={item.stock ? "green" : "red"}
                      variant="subtle"
                    >
                      {item.stock ? "在庫あり" : "在庫なし"}
                    </Badge>
                  </Table.Cell>
                  <Table.Cell>
                    <Text lineClamp={1} maxW="xs">
                      {item.description}
                    </Text>
                  </Table.Cell>
                  <Table.Cell>
                    <Text fontSize="sm" color="gray.500">
                      {new Date(item.created_at).toLocaleDateString("ja-JP")}
                    </Text>
                  </Table.Cell>
                  <Table.Cell onClick={(e) => e.stopPropagation()}>
                    <HStack gap={2}>
                      <IconButton
                        aria-label="編集"
                        size="sm"
                        variant="ghost"
                        onClick={() => handleEdit(item)}
                      >
                        <FiEdit color="blue" />
                      </IconButton>
                      <IconButton
                        aria-label="削除"
                        size="sm"
                        variant="ghost"
                        onClick={() => handleDeleteClick(item.item_id)}
                      >
                        <FiTrash2 color="red" />
                      </IconButton>
                    </HStack>
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table.Root>

          {items.length === 0 && (
            <Box textAlign="center" py={8}>
              <Text color="gray.500">商品が登録されていません</Text>
            </Box>
          )}
        </Box>

        {/* ページネーション */}
        {totalPages > 1 && (
          <Flex justify="center">
            <PaginationRoot
              count={items.length}
              pageSize={itemsPerPage}
              page={currentPage}
              onPageChange={(e) => setCurrentPage(e.page)}
            >
              <HStack>
                <PaginationPrevTrigger />
                <PaginationItems />
                <PaginationNextTrigger />
              </HStack>
            </PaginationRoot>
          </Flex>
        )}
      </VStack>

      {/* 詳細モーダル */}
      <DialogRoot
        open={isDetailModalOpen}
        onOpenChange={(e) => setIsDetailModalOpen(e.open)}
        size="xl"
      >
        <DialogContent>
          <DialogHeader>
            <DialogTitle>商品詳細</DialogTitle>
          </DialogHeader>
          <DialogCloseTrigger />
          <DialogBody>
            {selectedItem && <AdminItemDetail itemId={selectedItem.item_id} />}
          </DialogBody>
        </DialogContent>
      </DialogRoot>

      {/* 編集モーダル */}
      <DialogRoot
        open={isEditModalOpen}
        onOpenChange={(e) => setIsEditModalOpen(e.open)}
        size="xl"
      >
        <DialogContent>
          <DialogHeader>
            <DialogTitle>商品編集</DialogTitle>
          </DialogHeader>
          <DialogCloseTrigger />
          <DialogBody pb={6}>
            {selectedItem && (
              <AdminItemForm
                mode="edit"
                itemId={selectedItem.item_id}
                onSuccess={() => {
                  setIsEditModalOpen(false);
                  fetchItems();
                }}
                onCancel={() => setIsEditModalOpen(false)}
              />
            )}
          </DialogBody>
        </DialogContent>
      </DialogRoot>

      {/* 削除確認ダイアログ */}
      <DialogRoot
        open={isDeleteAlertOpen}
        onOpenChange={(e) => setIsDeleteAlertOpen(e.open)}
      >
        <DialogContent>
          <DialogHeader>
            <DialogTitle>商品の削除</DialogTitle>
          </DialogHeader>
          <DialogCloseTrigger />
          <DialogBody>
            この商品を削除してもよろしいですか？この操作は取り消せません。
          </DialogBody>
          <DialogFooter>
            <DialogActionTrigger asChild>
              <Button variant="outline">キャンセル</Button>
            </DialogActionTrigger>
            <Button colorScheme="red" onClick={handleDelete}>
              削除
            </Button>
          </DialogFooter>
        </DialogContent>
      </DialogRoot>
    </Box>
  );
};

export default AdminItemList;
