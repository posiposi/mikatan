import { useEffect, useState } from "react";
import { Table } from "@chakra-ui/react";
import { Book } from "../../types/book";
import { Box, Spinner } from "@chakra-ui/react";
import BookEditButton from "./BookEditButton";
import BookDeleteButton from "./BookDeleteButton";
import { getBooks } from "../utils/getBooks";
import BookRegisterDialog from "../components/BookRegisterDialog";
import ProgressAndReviewCard from "@/components/ProgressAndReviewCard";
import progressPercentageContext from "@/components/contexts/progressPercentageContext";

const BookTable = () => {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const fetchBooks = async () => {
    try {
      const data = await getBooks();
      setBooks(data);
    } catch (error) {
      console.error("Error fetching books:", error);
      alert("書籍の取得に失敗しました。");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100vh"
      >
        <Spinner size="xl" />
      </Box>
    );
  }

  return (
    <>
      <div className="book_register_btn">
        <BookRegisterDialog fetchBooks={fetchBooks} />
      </div>
      <Table.Root size="sm">
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeader>タイトル</Table.ColumnHeader>
            <Table.ColumnHeader>著者</Table.ColumnHeader>
            <Table.ColumnHeader>ジャンル</Table.ColumnHeader>
            <Table.ColumnHeader>出版社</Table.ColumnHeader>
            <Table.ColumnHeader>出版年</Table.ColumnHeader>
            <Table.ColumnHeader>総ページ数</Table.ColumnHeader>
            <Table.ColumnHeader>現状ページ</Table.ColumnHeader>
            <Table.ColumnHeader>進捗率</Table.ColumnHeader>
            <Table.ColumnHeader>金額</Table.ColumnHeader>
            <Table.ColumnHeader></Table.ColumnHeader>
            <Table.ColumnHeader></Table.ColumnHeader>
            <Table.ColumnHeader textAlign="end"></Table.ColumnHeader>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {books.map((book) => (
            <Table.Row key={book.id}>
              <Table.Cell>{book.title}</Table.Cell>
              <Table.Cell>{book.author}</Table.Cell>
              <Table.Cell>{book.genre}</Table.Cell>
              <Table.Cell>{book.publisher}</Table.Cell>
              <Table.Cell>{book.publishedAt}年</Table.Cell>
              <Table.Cell>{book.totalPage}p</Table.Cell>
              <Table.Cell>{book.progressPage}p</Table.Cell>
              <Table.Cell>{book.progressPercentage}%</Table.Cell>
              <Table.Cell>¥{Number(book.price).toLocaleString()}</Table.Cell>
              <Table.Cell>
                <BookEditButton book={book} fetchBooks={fetchBooks} />
              </Table.Cell>
              <Table.Cell>
                <progressPercentageContext.Provider
                  value={book.progressPercentage ?? 0}
                >
                  <ProgressAndReviewCard />
                </progressPercentageContext.Provider>
              </Table.Cell>
              <Table.Cell textAlign="end">
                <BookDeleteButton
                  bookId={book.id}
                  fetchBooks={fetchBooks}
                ></BookDeleteButton>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </>
  );
};

export default BookTable;
