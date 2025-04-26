import { HStack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { AiFillDelete } from "react-icons/ai";
import {
  DialogActionTrigger,
  DialogBody,
  DialogCloseTrigger,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogRoot,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import React from "react";

interface BookDeleteButtonProps {
  bookId: string;
  fetchBooks: () => void;
}

const BookDeleteButton: React.FC<BookDeleteButtonProps> = ({
  bookId,
  fetchBooks,
}) => {
  const [isDialogOpen, setIsDialogOpen] = React.useState(false);
  const deleteBook = async () => {
    try {
      const baseURL = import.meta.env.VITE_API_BASE_URL;
      const response = await fetch(`${baseURL}/v1/books/${bookId}`, {
        method: "DELETE",
      });
      if (!response.ok) {
        throw new Error("書籍の削除に失敗しました。");
      }
      alert("書籍を削除しました。");
      setIsDialogOpen(false);
      fetchBooks();
    } catch (error: unknown) {
      console.log("Error deleting book:", error);
      alert("書籍の削除に失敗しました。");
    }
  };

  return (
    <HStack>
      <DialogRoot
        open={isDialogOpen}
        onOpenChange={(details: { open: boolean }) =>
          setIsDialogOpen(details.open)
        }
        role="alertdialog"
      >
        <DialogTrigger asChild>
          <Button variant="outline" size="sm">
            <AiFillDelete /> Delete
          </Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>書籍削除</DialogTitle>
          </DialogHeader>
          <DialogBody>
            <p>本当に削除してもよろしいですか？</p>
          </DialogBody>
          <DialogFooter>
            <DialogActionTrigger asChild>
              <Button variant="outline">Cancel</Button>
            </DialogActionTrigger>
            <Button color={"red"} onClick={() => deleteBook()}>
              Delete
            </Button>
          </DialogFooter>
          <DialogCloseTrigger />
        </DialogContent>
      </DialogRoot>
    </HStack>
  );
};

export default BookDeleteButton;
