import { useEffect, useState } from "react";
import ItemCard from "./ItemCard";
import { Box, Spinner, Text, SimpleGrid } from "@chakra-ui/react";
import { getItems } from "../utils/getItems";
import { Item } from "../../types/item";

export default function TopPage() {
  const [items, setItems] = useState<Item[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchItems = async () => {
      try {
        const data = await getItems();
        setItems(data);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load items");
      } finally {
        setLoading(false);
      }
    };

    fetchItems();
  }, []);

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="50vh"
      >
        <Spinner size="xl" />
      </Box>
    );
  }

  if (error) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="50vh"
      >
        <Text color="red.500">Error: {error}</Text>
      </Box>
    );
  }

  if (items.length === 0) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="50vh"
      >
        <Text>No items available</Text>
      </Box>
    );
  }

  return (
    <Box p={4} pt={20}>
      <Text
        fontSize={{ base: "3xl", md: "4xl", lg: "5xl" }}
        fontWeight="extrabold"
        mb={8}
        textAlign="center"
        color="blue.600"
        letterSpacing="tight"
        lineHeight="shorter"
        textShadow="2px 2px 4px rgba(0,0,0,0.1)"
      >
        みかたんにってぃんぐ
      </Text>
      <SimpleGrid columns={{ base: 1, md: 2, lg: 3 }} gap={6}>
        {items.map((item) => (
          <ItemCard key={item.item_id} item={item} />
        ))}
      </SimpleGrid>
    </Box>
  );
}
