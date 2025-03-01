"use client"
import { zodResolver } from "@hookform/resolvers/zod";
import { Label } from "@radix-ui/react-label";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { register } from "module";
import { useSearchParams } from "next/navigation";

const resetPasswordSchema = z
  .object({
    new_password: z.string().min(6),
    confirm_password: z.string().min(6),
  })
  .refine((data) => data.new_password === data.confirm_password, {
    message: "Passwords do not match",
    path: ["confirm_password"],
  });

type FormData = z.infer<typeof resetPasswordSchema>;

export function ResetPasswordForm() {
    const token = useSearchParams().get("token")
  const {
    handleSubmit,
    register,
    formState: { isSubmitting, errors },
  } = useForm<FormData>({
    resolver: zodResolver(resetPasswordSchema),
  });

  const onSubmit = (data: FormData) => {
    console.log(data);
  };
  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <Label htmlFor="new_password">New Password</Label>
        <Input
          id="new_password"
          type="password"
          {...register("new_password")}
          error={errors.new_password?.message}
        />
      </div>
      <div>
        <Label htmlFor="confirm_password">Confirm Password</Label>
        <Input
          id="confirm_password"
          type="password"
          {...register("confirm_password")}
          error={errors.confirm_password?.message}
        />
      </div>
      <Button type="submit" className="w-full" loading={isSubmitting}>
        Reset Password
      </Button>
    </form>
  );
}
