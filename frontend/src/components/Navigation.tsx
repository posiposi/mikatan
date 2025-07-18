import { Box, Button, Flex, Heading, Spacer } from "@chakra-ui/react";
import { Link } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export default function Navigation() {
  const { isAuthenticated, logout } = useAuth();

  const handleLogout = async () => {
    try {
      await fetch("/v1/logout", {
        method: "POST",
        credentials: "include",
      });
      logout();
    } catch {
      logout();
    }
  };

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
          {!isAuthenticated && (
            <>
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
            </>
          )}
          {isAuthenticated && (
            <Button
              colorScheme="blue"
              variant="outline"
              size="sm"
              onClick={handleLogout}
            >
              ログアウト
            </Button>
          )}
        </Flex>
      </Flex>
    </Box>
  );
}
