import {
  DialogActionTrigger,
  DialogBody,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogRoot,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { useRef } from "react";
import { Button, Input, Stack } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import { Book } from "../../types/book";
import { SubmitHandler, useForm } from "react-hook-form";

interface BookEditButtonProps {
  book: Book;
  fetchBooks: () => void;
}

// TODO typeを共通化する
interface BookInput {
  id: string;
  title: string;
  genre: string;
  totalPage: string;
  progressPage: string;
  author: string;
  publisher: string;
  publishedAt: string;
  price: string;
}

const BookEditButton: React.FC<BookEditButtonProps> = ({
  book,
  fetchBooks,
}) => {
  const { register, handleSubmit, reset } = useForm<BookInput>();
  // 編集ボタン押下時イベント定義
  const onEditSubmit: SubmitHandler<BookInput> = async (data) => {
    // TODO 登録処理と同様のため共通化する
    const publishedAtYear = new Date(data.publishedAt).getFullYear();
    const totalPage = parseInt(data.totalPage);
    const progressPage = parseInt(data.progressPage);
    const price = parseInt(data.price);
    const bookData: Book = {
      ...data,
      publishedAt: publishedAtYear,
      totalPage: totalPage,
      progressPage: progressPage,
      price: price,
    };

    try {
      const baseURL = import.meta.env.VITE_API_BASE_URL;
      const response = await fetch(`${baseURL}/v1/books/${book.id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(bookData),
      });
      if (!response.ok) {
        throw new Error("書籍の編集登録に失敗しました。");
      }
      alert("書籍情報を変更しました。");
      reset();
      fetchBooks();
    } catch (error: unknown) {
      alert(error);
    }
  };

  const ref = useRef<HTMLInputElement>(null);
  return (
    <DialogRoot initialFocusEl={() => ref.current}>
      <DialogTrigger asChild>
        <Button variant="outline">Edit</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>書籍編集</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit(onEditSubmit)}>
          <DialogBody pb="4">
            <Stack gap="4">
              <Field label="タイトル">
                <Input defaultValue={book.title} {...register("title")} />
              </Field>
              <Field label="著者">
                <Input defaultValue={book.author} {...register("author")} />
              </Field>
              <Field label="ジャンル">
                <Input defaultValue={book.genre} {...register("genre")} />
              </Field>
              <Field label="出版社">
                <Input
                  defaultValue={book.publisher}
                  {...register("publisher")}
                />
              </Field>
              <Field label="出版年">
                <Input
                  defaultValue={String(book.publishedAt)}
                  {...register("publishedAt")}
                />
              </Field>
              <Field label="総ページ数">
                <Input
                  defaultValue={String(book.totalPage)}
                  {...register("totalPage")}
                />
              </Field>
              <Field label="現状ページ">
                <Input
                  defaultValue={String(book.progressPage)}
                  {...register("progressPage")}
                />
              </Field>
              <Field label="金額">
                <Input
                  defaultValue={String(book.price)}
                  {...register("price")}
                />
              </Field>
            </Stack>
          </DialogBody>
          <DialogFooter>
            <DialogActionTrigger asChild>
              <Button variant="outline">Cancel</Button>
            </DialogActionTrigger>
            <DialogActionTrigger asChild>
              <Button type="submit" variant="outline" colorPalette="blue">
                更新
              </Button>
            </DialogActionTrigger>
          </DialogFooter>
        </form>
      </DialogContent>
    </DialogRoot>
  );
};

export default BookEditButton;
