import { Input, Stack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { AiFillBook } from "react-icons/ai";
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
import { Field } from "@/components/ui/field";
import { useForm, SubmitHandler } from "react-hook-form";
import { Book } from "../../types/book";

interface BookInput {
  id: string;
  title: string;
  genre: string;
  totalPage: string;
  author: string;
  publisher: string;
  publishedAt: string;
  price: string;
}

interface BookRegisterDialogProps {
  fetchBooks: () => void;
}

const BookRegisterDialog: React.FC<BookRegisterDialogProps> = ({
  fetchBooks,
}) => {
  const { register, handleSubmit, reset } = useForm<BookInput>();
  const onSubmit: SubmitHandler<BookInput> = async (data) => {
    const publishedAtYear = new Date(data.publishedAt).getFullYear();
    const totalPage = parseInt(data.totalPage);
    const price = parseInt(data.price);
    const bookData: Book = {
      ...data,
      publishedAt: publishedAtYear,
      totalPage: totalPage,
      price: price,
    };

    try {
      const baseURL = import.meta.env.VITE_API_BASE_URL;
      const response = await fetch(`${baseURL}/v1/books`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(bookData),
      });
      if (!response.ok) {
        throw new Error("書籍の登録に失敗しました。");
      }
      alert("書籍を登録しました。");
      reset();
      fetchBooks();
    } catch (error: unknown) {
      alert(error);
    }
  };
  return (
    <DialogRoot>
      <DialogTrigger asChild>
        <Button variant="outline">
          <AiFillBook /> Register
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>書籍登録</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit(onSubmit)}>
          <DialogBody pb="4">
            <Stack gap="4">
              <Field label="タイトル">
                <Input {...register("title")} />
              </Field>
              <Field label="著者">
                <Input {...register("author")} />
              </Field>
              <Field label="ジャンル">
                <Input {...register("genre")} />
              </Field>
              <Field label="出版社">
                <Input {...register("publisher")} />
              </Field>
              <Field label="出版年">
                <Input {...register("publishedAt")} />
              </Field>
              <Field label="総ページ数">
                <Input {...register("totalPage")} />
              </Field>
              <Field label="金額">
                <Input {...register("price")} />
              </Field>
            </Stack>
          </DialogBody>
          <DialogFooter>
            <DialogActionTrigger asChild>
              <Button variant="outline">Cancel</Button>
            </DialogActionTrigger>
            <DialogActionTrigger asChild>
              <Button type="submit" variant="outline" colorPalette="blue">
                Save
              </Button>
            </DialogActionTrigger>
          </DialogFooter>
        </form>
      </DialogContent>
    </DialogRoot>
  );
};

export default BookRegisterDialog;
