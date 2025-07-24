import React from "react";
import { Link } from "react-router-dom";
import { Box, Heading, Grid, Text, Button, VStack } from "@chakra-ui/react";
import { useColorModeValue } from "../hooks/color-mode-hooks";

const AdminHome: React.FC = () => {
  const cardBg = useColorModeValue("white", "gray.800");
  const textColor = useColorModeValue("gray.600", "gray.300");

  return (
    <Box minH="100vh" py={8}>
      <VStack gap={8} w="full" align="center" maxW="6xl" mx="auto" px={4}>
        <Heading as="h1" size="2xl" textAlign="center">
          管理画面ホーム
        </Heading>
        <Grid
          templateColumns={{ base: "1fr", md: "1fr 1fr" }}
          gap={6}
          w="full"
          maxW="4xl"
        >
          <Box bg={cardBg} p={6} borderRadius="lg" shadow="md">
            <Heading as="h2" size="lg" mb={4}>
              商品管理
            </Heading>
            <Text color={textColor} mb={4}>
              商品の一覧表示、編集、削除ができます。
            </Text>
            <Link to="/admin/items">
              <Button colorScheme="blue" size="md">
                商品一覧を見る
              </Button>
            </Link>
          </Box>
          <Box bg={cardBg} p={6} borderRadius="lg" shadow="md">
            <Heading as="h2" size="lg" mb={4}>
              新規商品登録
            </Heading>
            <Text color={textColor} mb={4}>
              新しい商品を登録できます。
            </Text>
            <Link to="/admin/items/new">
              <Button colorScheme="green" size="md">
                商品を登録する
              </Button>
            </Link>
          </Box>
        </Grid>
      </VStack>
    </Box>
  );
};

export default AdminHome;
