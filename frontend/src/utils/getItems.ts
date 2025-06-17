import { Item } from "../../types/item";

const baseURL = import.meta.env.VITE_API_BASE_URL || "http://localhost:8080";

export const getItems = async (): Promise<Item[]> => {
  try {
    const response = await fetch(`${baseURL}/v1/items`);
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
