import { Button, Fieldset, Input, Stack } from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import {
  NativeSelectField,
  NativeSelectRoot,
} from "@/components/ui/native-select";

interface BookEditDialogProps {
  onClose: () => void;
}

const BookEditDialog = ({ onClose }: BookEditDialogProps) => {
  return (
    <Fieldset.Root size="lg" maxW="md">
      <Stack>
        <Fieldset.Legend>Contact details</Fieldset.Legend>
        <Fieldset.HelperText>
          Please provide your contact details below.
        </Fieldset.HelperText>
      </Stack>
      <Fieldset.Content>
        <Field label="Name">
          <Input name="name" />
        </Field>
        <Field label="Email address">
          <Input name="email" type="email" />
        </Field>
        <Field label="Country">
          <NativeSelectRoot>
            <NativeSelectField
              name="country"
              items={[
                "United Kingdom (UK)",
                "Canada (CA)",
                "United States (US)",
              ]}
            />
          </NativeSelectRoot>
        </Field>
      </Fieldset.Content>
      <Button onClick={onClose}>Close</Button>
      <Button type="submit" alignSelf="flex-start">
        Submit
      </Button>
    </Fieldset.Root>
  );
};

export default BookEditDialog;
