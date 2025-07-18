import { Item } from "../../types/item";
import { get } from "./api";

export const getItems = async (): Promise<Item[]> => {
  try {
    const response = await get("/v1/items", true);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Failed to fetch items:", error);
    throw error;
  }
};
