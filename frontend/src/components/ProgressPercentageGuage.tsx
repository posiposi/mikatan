import { HStack, Progress } from "@chakra-ui/react";
import { useContext } from "react";
import progressPercentageContext from "@/components/contexts/progressPercentageContext";

const ProgressPercentageGuage = () => {
  const progressPercentage = useContext(progressPercentageContext);
  return (
    <Progress.Root defaultValue={progressPercentage} maxW="sm">
      <HStack gap="5">
        <Progress.Label>進捗率</Progress.Label>
        <Progress.Track flex="1">
          <Progress.Range />
        </Progress.Track>
        <Progress.ValueText>{progressPercentage}%</Progress.ValueText>
      </HStack>
    </Progress.Root>
  );
};

export default ProgressPercentageGuage;
