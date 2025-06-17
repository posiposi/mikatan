import {
  Button,
  Card,
  Image,
  Badge,
  HStack,
  Text,
  Box,
  VStack,
  Tooltip,
  Circle,
} from "@chakra-ui/react";
import { LuCheck, LuX } from "react-icons/lu";
import { Item } from "../../types/item";

interface ItemCardProps {
  item: Item;
}

const ItemCard = ({ item }: ItemCardProps) => {
  return (
    <Card.Root
      maxW="sm"
      overflow="hidden"
      transition="all 0.3s"
      _hover={{ transform: "translateY(-4px)", shadow: "lg" }}
    >
      <Box position="relative">
        <Image
          src="https://images.unsplash.com/photo-1555041469-a586c61ea9bc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80"
          alt={item.item_name}
          opacity={item.stock ? 1 : 0.6}
          transition="opacity 0.3s"
        />
        {!item.stock && (
          <Box
            position="absolute"
            top="50%"
            left="50%"
            transform="translate(-50%, -50%)"
            bg="blackAlpha.700"
            color="white"
            px={4}
            py={2}
            borderRadius="md"
            fontWeight="bold"
            fontSize="lg"
          >
            売り切れ
          </Box>
        )}
      </Box>
      <Card.Body gap="3">
        <VStack align="stretch" gap={3}>
          <Card.Title>{item.item_name}</Card.Title>
          <Card.Description>
            {item.description || "No description available"}
          </Card.Description>

          <HStack gap={2} align="center">
            <Circle
              size="6"
              bg={item.stock ? "green.100" : "red.100"}
              color={item.stock ? "green.600" : "red.600"}
            >
              {item.stock ? <LuCheck size={14} /> : <LuX size={14} />}
            </Circle>
            <Badge
              colorPalette={item.stock ? "green" : "red"}
              variant={item.stock ? "subtle" : "solid"}
              px={3}
              py={1}
              borderRadius="full"
              fontSize="sm"
            >
              {item.stock ? "在庫あり" : "在庫なし"}
            </Badge>
            {item.stock && (
              <Tooltip.Root>
                <Tooltip.Trigger asChild>
                  <Text fontSize="xs" color="gray.500">
                    ✓ 即日発送
                  </Text>
                </Tooltip.Trigger>
                <Tooltip.Positioner>
                  <Tooltip.Content>すぐに発送可能</Tooltip.Content>
                </Tooltip.Positioner>
              </Tooltip.Root>
            )}
          </HStack>
        </VStack>
      </Card.Body>
      <Card.Footer gap="2">
        <Button
          variant="solid"
          color={"white"}
          colorPalette={item.stock ? "orange" : "gray"}
          disabled={!item.stock}
          size="sm"
          flex={1}
          _disabled={{
            opacity: 0.6,
            cursor: "not-allowed",
            _hover: { bg: "gray.300" },
          }}
        >
          {item.stock ? "今すぐ購入" : "入荷待ち"}
        </Button>
        <Button
          variant="outline"
          colorPalette={item.stock ? "blue" : "gray"}
          disabled={!item.stock}
          size="sm"
          flex={1}
          _disabled={{
            opacity: 0.6,
            cursor: "not-allowed",
          }}
        >
          カートに追加
        </Button>
      </Card.Footer>
    </Card.Root>
  );
};

export default ItemCard;
