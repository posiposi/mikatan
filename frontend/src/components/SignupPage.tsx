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

interface SignupFormData {
  name: string;
  email: string;
  password: string;
}

export default function SignupPage() {
  const [isLoading, setIsLoading] = useState(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignupFormData>();
  const navigate = useNavigate();
  const { login } = useAuth();

  const onSubmit = async (data: SignupFormData) => {
    setIsLoading(true);
    try {
      const signupResponse = await post("/v1/signup", data);

      if (signupResponse.ok) {
        const loginResponse = await post("/v1/login", {
          email: data.email,
          password: data.password,
        });

        if (loginResponse.ok) {
          const result = await loginResponse.json();
          login(result.token);
          alert("会員登録が完了しました。自動的にログインしています。");
          navigate("/");
        } else {
          alert("会員登録は完了しましたが、ログインに失敗しました。手動でログインしてください。");
          navigate("/login");
        }
      } else {
        const error = await signupResponse.text();
        alert(`登録に失敗しました: ${error}`);
      }
    } catch {
      alert("ネットワークエラーが発生しました。");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Container py={8} pt={20}>
      <Box
        className="register_form"
        p={8}
        borderRadius="lg"
        boxShadow="lg"
        width="800px"
      >
        <Heading mb={6} textAlign="center">
          会員登録
        </Heading>
        <form onSubmit={handleSubmit(onSubmit)}>
          <VStack gap={4}>
            <Field
              label="名前"
              invalid={!!errors.name}
              errorText={errors.name?.message}
            >
              <Input
                {...register("name", { required: "名前は必須です" })}
                placeholder="名前を入力してください"
                autoComplete="name"
              />
            </Field>

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
                autoComplete="new-password"
              />
            </Field>

            <Button
              type="submit"
              colorScheme="blue"
              width="full"
              loading={isLoading}
              loadingText="登録中..."
            >
              会員登録
            </Button>
          </VStack>
        </form>
      </Box>
    </Container>
  );
}
