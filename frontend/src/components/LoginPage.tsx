import { useState } from "react";
import {
  Box,
  Button,
  Input,
  VStack,
  Heading,
  Container,
} from "@chakra-ui/react";
import { Field } from "@/components/ui/field";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { post } from "../utils/api";

interface LoginFormData {
  email: string;
  password: string;
}

export default function LoginPage() {
  const [isLoading, setIsLoading] = useState(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>();
  const navigate = useNavigate();
  const { login } = useAuth();

  const onSubmit = async (data: LoginFormData) => {
    setIsLoading(true);
    try {
      const response = await post("/v1/login", data);

      if (response.ok) {
        const result = await response.json();
        login(result.token);
        alert("ログインしました。");
        navigate("/");
      } else {
        const error = await response.text();
        alert(`ログインに失敗しました: ${error}`);
      }
    } catch {
      alert("ネットワークエラーが発生しました");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Container py={8} pt={20}>
      <Box
        className="login_form"
        p={8}
        borderRadius="lg"
        boxShadow="lg"
        width="800px"
      >
        <Heading mb={6} textAlign="center">
          ログイン
        </Heading>
        <form onSubmit={handleSubmit(onSubmit)}>
          <VStack gap={4}>
            <Field
              label="メールアドレス"
              invalid={!!errors.email}
              errorText={errors.email?.message}
            >
              <Input
                type="email"
                {...register("email", {
                  required: "メールアドレスは必須です",
                  pattern: {
                    value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
                    message: "有効なメールアドレスを入力してください",
                  },
                })}
                placeholder="email@example.com"
                autoComplete="email"
              />
            </Field>

            <Field
              label="パスワード"
              invalid={!!errors.password}
              errorText={errors.password?.message}
            >
              <Input
                type="password"
                {...register("password", {
                  required: "パスワードは必須です",
                  minLength: {
                    value: 6,
                    message: "パスワードは6文字以上で入力してください",
                  },
                })}
                placeholder="パスワードを入力してください"
                autoComplete="current-password"
              />
            </Field>

            <Button
              type="submit"
              colorScheme="blue"
              width="full"
              loading={isLoading}
              loadingText="ログイン中..."
            >
              ログイン
            </Button>
          </VStack>
        </form>
      </Box>
    </Container>
  );
}
