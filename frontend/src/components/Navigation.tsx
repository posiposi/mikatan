import { Box, Button, Flex, Heading, Spacer } from "@chakra-ui/react";
import { Link } from "react-router-dom";
import { useState } from "react";
import { useAuth } from "../hooks/useAuth";
import LogoutConfirmDialog from "./LogoutConfirmDialog";
import { post } from "../utils/api";

export default function Navigation() {
  const { isAuthenticated, logout } = useAuth();
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const handleLogout = async () => {
    try {
      await post("/v1/logout");
      logout();
    } catch {
      logout();
    }
  };

  const handleConfirmLogout = () => {
    setIsDialogOpen(false);
    handleLogout();
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
              onClick={() => setIsDialogOpen(true)}
            >
              ログアウト
            </Button>
          )}
        </Flex>
      </Flex>
      <LogoutConfirmDialog
        isOpen={isDialogOpen}
        onClose={() => setIsDialogOpen(false)}
        onConfirm={handleConfirmLogout}
      />
    </Box>
  );
}
