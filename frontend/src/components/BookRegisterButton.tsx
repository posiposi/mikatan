import { HStack } from "@chakra-ui/react";
import { Button } from "@/components/ui/button";
import { AiFillBook } from "react-icons/ai";

const BookRegisterButton = () => {
  return (
    <>
      <HStack>
        <Button colorPalette="teal" variant="solid">
          <AiFillBook /> Register
        </Button>
      </HStack>
    </>
  );
};

export default BookRegisterButton;
