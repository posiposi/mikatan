export const getBooks = async () => {
  const baseURL = import.meta.env.VITE_API_BASE_URL;
  const response = await fetch(`${baseURL}/v1/books`);
  if (!response.ok) {
    throw new Error("書籍の取得に失敗しました。");
  }
  const data = await response.json();
  return data;
};
