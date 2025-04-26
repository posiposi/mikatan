import ItemCard from "./ItemCard";
import { Box } from "@chakra-ui/react";

export default function TopPage() {
  return (
    <Box display="flex" justifyContent="space-around">
      <ItemCard />
      <ItemCard />
      <ItemCard />
    </Box>
  );
}
