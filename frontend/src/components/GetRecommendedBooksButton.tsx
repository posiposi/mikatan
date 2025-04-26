import { Button, Spinner, Center } from "@chakra-ui/react";
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
import { AiFillBook } from "react-icons/ai";
import { useState } from "react";

const GetRecommendBooksButton = () => {
  const [loading, setLoading] = useState<boolean>(false);
  const [responseMessage, setResponseMessage] = useState<string>("");

  const handleRecommend = async () => {
    setLoading(true);
    try {
      const baseURL = import.meta.env.VITE_API_BASE_URL;
      const response = await fetch(`${baseURL}/v1/books/recommend`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (!response.ok) {
        throw new Error("error occurred");
      }
      const data = await response.json();
      setResponseMessage(data.content);
      setLoading(false);
    } catch (error: unknown) {
      console.error(error);
      setResponseMessage("error occurred");
      setLoading(false);
    }
  };

  const handleClose = () => {
    setResponseMessage("");
  };

  return (
    <>
      {loading && (
        <Center
          position="fixed"
          top="0"
          left="0"
          width="100vw"
          height="100vh"
          backgroundColor="rgba(0, 0, 0, 0.3)"
        >
          <Spinner size="xl" />
        </Center>
      )}
      <DialogRoot onEscapeKeyDown={handleClose}>
        <DialogTrigger asChild>
          <Button variant="outline" size="sm" onClick={handleRecommend}>
            {loading ? (
              <Spinner size="sm" />
            ) : (
              <>
                <AiFillBook /> Recommend
              </>
            )}
          </Button>
        </DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>次に読むおすすめ書籍は…</DialogTitle>
          </DialogHeader>
          <DialogBody>
            <p>{responseMessage}</p>
          </DialogBody>
          <DialogFooter>
            <DialogActionTrigger asChild>
              <Button variant="outline" onClick={handleClose}>
                Close
              </Button>
            </DialogActionTrigger>
          </DialogFooter>
        </DialogContent>
      </DialogRoot>
    </>
  );
};

export default GetRecommendBooksButton;
