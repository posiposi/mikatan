import React from "react";
import { Link, Outlet } from "react-router-dom";
import {
  Box,
  Flex,
  VStack,
  Text,
  Container,
} from "@chakra-ui/react";
import { FiHome, FiList, FiPlus } from "react-icons/fi";
import { useAuth } from "../../hooks/useAuth";
import { useColorModeValue } from "../hooks/color-mode-hooks";

interface SidebarProps {
  onClose?: () => void;
}

const SidebarContent: React.FC<SidebarProps> = ({ onClose }) => {
  const bg = useColorModeValue("gray.50", "gray.900");
  const borderColor = useColorModeValue("gray.200", "gray.700");
  const hoverBg = useColorModeValue("cyan.400", "cyan.600");

  const menuItems = [
    { to: "/admin", label: "ダッシュボード", icon: FiHome },
    { to: "/admin/items", label: "商品一覧", icon: FiList },
    { to: "/admin/items/new", label: "商品登録", icon: FiPlus },
  ];

  return (
    <Box
      bg={bg}
      borderRight="1px"
      borderRightColor={borderColor}
      w={{ base: "full", md: 60 }}
      pos="fixed"
      top="0"
      left="0"
      h="100vh"
      zIndex="1000"
    >
      <Flex h="20" alignItems="center" mx="8">
        <Text fontSize="2xl" fontFamily="monospace" fontWeight="bold">
          管理画面
        </Text>
      </Flex>
      <VStack align="stretch" gap={1} px="4">
        {menuItems.map((item) => (
          <Box key={item.to}>
            <Link to={item.to} onClick={onClose}>
              <Flex
                align="center"
                p="4"
                mx="4"
                borderRadius="lg"
                role="group"
                cursor="pointer"
                _hover={{
                  bg: hoverBg,
                  color: "white",
                }}
              >
                <Box mr="4">
                  <item.icon size={16} />
                </Box>
                {item.label}
              </Flex>
            </Link>
          </Box>
        ))}
      </VStack>
    </Box>
  );
};

const AdminDashboard: React.FC = () => {
  const { isAdmin } = useAuth();
  const bgColor = useColorModeValue("gray.100", "gray.900");

  if (!isAdmin) {
    return (
      <Container centerContent>
        <Box textAlign="center" py={10} px={6}>
          <Text fontSize="2xl" fontWeight="bold" color="red.500" mb={4}>
            アクセス権限がありません
          </Text>
          <Text color="gray.500">管理者権限が必要です。</Text>
        </Box>
      </Container>
    );
  }

  return (
    <Box minH="100vh" bg={bgColor}>
      <SidebarContent />
      <Box ml={{ base: 0, md: 60 }} p="4">
        <Box p="4">
          <Outlet />
        </Box>
      </Box>
    </Box>
  );
};

export default AdminDashboard;
