import { Box, Button, Flex, Heading, Spacer } from "@chakra-ui/react";
import { Link } from "react-router-dom";

export default function Navigation() {
  return (
    <Box
      bg="blue.500"
      px={4}
      py={2}
      position="fixed"
      top={0}
      left={0}
      right={0}
      zIndex={1000}
    >
      <Flex alignItems="center">
        <Heading size="md" color="white" fontWeight="bold">
          <Link
            to="/"
            style={{
              textDecoration: "none",
              color: "inherit",
              textShadow: "1px 1px 2px rgba(0,0,0,0.3)",
            }}
          >
            みかたんにってぃんぐ
          </Link>
        </Heading>
        <Spacer />
        <Flex gap={2}>
          <Link to="/login">
            <Button colorScheme="blue" variant="outline" size="sm">
              ログイン
            </Button>
          </Link>
          <Link to="/signup">
            <Button colorScheme="blue" variant="outline" size="sm">
              会員登録
            </Button>
          </Link>
        </Flex>
      </Flex>
    </Box>
  );
}
